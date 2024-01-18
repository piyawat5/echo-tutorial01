package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/piyawat5/echo-tutorial01/routes/movie"
)

func main(){
	e :=echo.New()

	e.GET("/movies", movie.GetAllMoviesHandler)
	e.GET("/movies/:id", movie.GetMoviebyId)
	e.POST("/movies", movie.PostMovie)




	port := "3100"
	log.Fatal(e.Start(":" + port))

}