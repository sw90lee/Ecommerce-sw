package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swlee/Ecommerce-sw/controllers"
	"github.com/swlee/Ecommerce-sw/database"
	"github.com/swlee/Ecommerce-sw/middleware"
	"github.com/swlee/Ecommerce-sw/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
