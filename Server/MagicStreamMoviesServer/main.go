package main

import (
	"context"
	"fmt"
	"log"

	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	//This is the main function
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello MagicStreamMovies")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to find .env file")
	}

	var client *mongo.Client = database.Connect()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	routes.SetupUnProtectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
