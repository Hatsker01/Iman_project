package api

import (
	//v1 "github.com/Hatsker01/Iman_project/api_getaway/api/handlers/v1"
	"github.com/Hatsker01/Iman_project/api_getaway/config"
	"github.com/Hatsker01/Iman_project/api_getaway/pkg/logger"
	services "github.com/Hatsker01/Iman_project/api_getaway/services"

	"github.com/gin-gonic/gin"

	_ "github.com/Hatsker01/Iman_project/api_getaway/api/docs"
	v1 "github.com/Hatsker01/Iman_project/api_getaway/api/handlers/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager

}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,

	})
	api := router.Group("/v1")
	api.GET("/users/:id", handlerV1.GetPost)
	api.PUT("/users/update", handlerV1.UpdatePost)
	api.GET("/users/allpost", handlerV1.GetAllPost)
	api.DELETE("/users/delete/:id",handlerV1.DeletePost)
	api.POST("/posts",handlerV1.CreatePost)
	
	//	api.GET("/users/lala/:first_name",handlerV1.SearchUser)
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//api.POST("/post",handlerV1.CreatePost)

	return router
}
