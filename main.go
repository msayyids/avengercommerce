package main

import (
	"echo/config"
	"echo/controller"
	"echo/middleware"
	"echo/repository"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDb()
	r := repository.Repo{DB: db}
	cn := controller.Controller{Repo: r}

	e := echo.New()
	e.POST("/register", cn.Register)
	e.POST("/login", cn.Login)

	ar := e.Group("/avenger")

	ar.GET("/products", middleware.AuthMiddleware(cn.GetAllProduct, db))
	ar.POST("/transaction", middleware.AuthMiddleware(cn.Transaction, db))
	ar.GET("/store", middleware.AuthMiddleware(cn.GetAllStore, db))
	ar.GET("/store/:id", middleware.AuthMiddleware(cn.StoreWeatherbyId, db))

	e.Logger.Fatal(e.Start(":8080"))
}
