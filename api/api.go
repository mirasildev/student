package api

import (
	v1 "github.com/mirasildev/student/api/v1"
	"github.com/mirasildev/student/config"
	"github.com/mirasildev/student/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/mirasildev/student/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

/// @title           Swagger for book api
// @version         1.0
// @description     This is a book service api.
// @host      localhost:8000
// @BasePath  /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	router.Static("/media", "./media")

	apiV1 := router.Group("/v1")

	// apiV1.GET("/student/", handlerV1.GetAllStudents)
	apiV1.POST("/student", handlerV1.CreateUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}