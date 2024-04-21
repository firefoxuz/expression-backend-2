package main

import (
	"context"
	"expression-backend/api/route"
	"expression-backend/internal/entities"
	"expression-backend/internal/services"
	agent "expression-backend/proto"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const serverAddr = "0.0.0.0:8085"

func initConfig() {
	viper.NewWithOptions()
	viper.SetConfigName(".env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	initConfig()

	r := route.NewRouterMux()
	r.RegisterRoutes()

	srv := &http.Server{
		Handler:           r.GetRouter(),
		Addr:              serverAddr,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Printf("Server started at %s", serverAddr)

	go func() {
		host := "localhost"
		port := "2552"

		addr := fmt.Sprintf("%s:%s", host, port) // используем адрес сервера
		// установим соединение
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Println("could not connect to grpc server: ", err)
			os.Exit(1)
		}
		// закроем соединение, когда выйдем из функции
		defer conn.Close()

		grpcClient := agent.NewMiniTaskServiceClient(conn)

		for {

			time.Sleep(3 * time.Second)

			expressions, err := entities.GetNotFinishedExpressions()
			if err != nil {
				log.Println(err)
				continue
			}

			//fmt.Println(solve, err)

			for _, expression := range *expressions {

				exp := services.NewExpression(expression.Expression)
				tokensAsString, err := exp.GetTokensAsString()
				if err != nil {
					expression.IsValid = false
					err = entities.UpdateExpression(expression)
					log.Println(err)
					continue
				}

				for {
					singleExps, _ := services.FindSingleExpressions(tokensAsString)

					for _, singleExp := range singleExps {
						solve, _ := grpcClient.Solve(context.Background(), &agent.MiniTaskRequest{
							ExpressionId: 1,
							Task:         singleExp,
						})
						tokensAsString = strings.ReplaceAll(tokensAsString, " "+singleExp+" ", " "+strconv.Itoa(int(solve.Result))+" ")
					}
					if len(singleExps) == 0 {
						break
					}
				}
				expression.Result = tokensAsString
				expression.IsFinished = true
				_ = entities.UpdateExpression(expression)
			}
		}

	}()

	srv.ListenAndServe()
}
