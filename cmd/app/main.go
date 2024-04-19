package main

import (
	"expression-backend/api/route"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
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

	srv.ListenAndServe()
}
