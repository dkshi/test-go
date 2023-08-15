package main

import (
	"fmt"
	"log"

	"github.com/dkshi/test-go/pkg/handler"
	"github.com/dkshi/test-go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	service := service.NewService()
	handler := handler.NewHandler(service)

	mux := handler.InitRoutes()
	mux.Logger.Fatal(mux.Start(fmt.Sprintf(":%s", viper.GetString("port"))))

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
