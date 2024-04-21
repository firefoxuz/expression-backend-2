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
	log.SetOutput(os.Stdout)

	log.Printf("Server started at %s", serverAddr)

	go func() {
		for {
		connectionStage:
			host := "agent"
			port := "2552"

			addr := fmt.Sprintf("%s:%s", host, port) // используем адрес сервера
			// установим соединение
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

			if err != nil {
				log.Println("could not connect to grpc server: ", err)
				time.Sleep(time.Second * 5)
				break
			}
			// закроем соединение, когда выйдем из функции
			defer conn.Close()

			grpcClient := agent.NewMiniTaskServiceClient(conn)

			for {

				time.Sleep(3 * time.Second)

				expressions, err := entities.GetNotFinishedExpressions()
				log.Println("expressions ->", expressions)
				if err != nil {
					log.Println(err)
					continue
				}

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
						log.Println("singleExps ->", singleExps)
						for _, singleExp := range singleExps {
							solve, err := grpcClient.Solve(context.Background(), &agent.MiniTaskRequest{
								ExpressionId: 1,
								Task:         singleExp,
							})
							log.Println("solve ->", solve, err)
							if err != nil {
								goto connectionStage
							}

							if solve.IsValid {
								tokensAsString = strings.ReplaceAll(tokensAsString, " "+singleExp+" ", " "+strconv.Itoa(int(solve.Result))+" ")
							} else {
								expression.IsFinished = true
								expression.IsValid = false
								_ = entities.UpdateExpression(expression)
							}
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
		}
	}()

	srv.ListenAndServe()
}
