package main

import (
	"siki/configs"
	"siki/controllers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	configs.SetupDB()

	server := echo.New()

	server.POST("/users", controllers.Create)
	server.GET("/users", controllers.Read)
	server.GET("/users/:id", controllers.Read)
	server.PUT("/users/:id", controllers.Update)
	server.DELETE("/users/:id", controllers.Delete)
	server.Start(":1323")

}
