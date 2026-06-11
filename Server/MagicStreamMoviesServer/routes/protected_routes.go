package routes

import (
	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movies/:imdb_id", controllers.GetMovie(client))
	router.POST("/addmovie", controllers.AddMovie(client))
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies(client))
	router.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate(client))
}
