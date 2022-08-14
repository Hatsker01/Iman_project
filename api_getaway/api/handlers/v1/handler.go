package v1

import (
	"github.com/Hatsker01/Iman_project/api_getaway/config"
	"github.com/Hatsker01/Iman_project/api_getaway/pkg/logger"
	services "github.com/Hatsker01/Iman_project/api_getaway/services"
	"github.com/Hatsker01/Iman_project/api_getaway/storage/repo"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redisStorage   repo.RepositoryStorage
}

type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.RepositoryStorage
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage:   c.Redis,
	}
}
