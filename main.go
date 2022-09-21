package main

import (
	"fmt"
	"log"
	"os"

	controller "github.com/ddld93/abedmis/src/controllers"
	"github.com/ddld93/abedmis/src/middleware"
	"github.com/ddld93/abedmis/src/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/ddld93/abedmis/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file >>> ", err)
// 	}

// }

// @title           Swagger Example API

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "database"
	}
	if port == "" {
		port = "3000"
	}
	BenenCtrl := controller.ConnectDB(host, 27017)
	route := routes.BeneRoute{BenenCtrl: BenenCtrl}
	router := gin.New()
	//router.Use(middleware.Logger())
	v1 := router.Group("/api/v3")
	v1.Use(middleware.Auth())
		{
			v1.POST("/", route.AddOne())
			v1.POST("/batch", route.AddMany())
			v1.GET("/", route.GetAll())
			v1.GET("/:id", route.GetOne())
			// v1.POST("/post", route.FormTest())
		}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(fmt.Sprintf("%s:%v", "0.0.0.0", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server running on port %s", port)
}
//...
