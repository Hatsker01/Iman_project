package main

import (

	"github.com/Hatsker01/Iman_project/api_getaway/api"
	"github.com/Hatsker01/Iman_project/api_getaway/config"
	"github.com/Hatsker01/Iman_project/api_getaway/pkg/logger"
	services "github.com/Hatsker01/Iman_project/api_getaway/services"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
