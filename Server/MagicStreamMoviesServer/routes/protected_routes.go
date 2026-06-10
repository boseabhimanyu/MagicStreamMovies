package routes

import (
	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/boseabhimanyu/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())
	router.GET("/movies/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies())
}
