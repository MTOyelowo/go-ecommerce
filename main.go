package main

import (
	"log"
	"os"

	"github.com/MTOyelowo/go-ecommerce/controllers"
	"github.com/MTOyelowo/go-ecommerce/database"
	"github.com/MTOyelowo/go-ecommerce/middleware"
	"github.com/MTOyelowo/go-ecommerce/routes"
	"github.com/gin-gonic/gin"
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

	router.Get("/addtocart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartcheckout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
