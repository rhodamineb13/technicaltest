package server

import (
	"technicaltest/config"
	"technicaltest/database"
	"technicaltest/handler"
	"technicaltest/repository"
	"technicaltest/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.NewConfig()
	database.ConnectDB()

	r := gin.Default()
	handler := handler.NewHandler(service.NewService(repository.NewRepository(database.DB)))

	r.POST("/create-client", handler.Create)
	r.GET("", handler.GetAll)
	r.GET("/:id", handler.GetByID)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)

	r.Run()
}
