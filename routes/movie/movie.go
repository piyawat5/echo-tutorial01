package movie

import (
	"strconv"

	"github.com/labstack/echo/v4"
)
type Movie struct {
	ImdbID string `json:"imdbID"`
	Title string `json:"title"`
	Year int `json:"year"`
	Rating float64 `json:"rating"`
}

 var movies =[]Movie{{
	ImdbID: "qwer", 
	Title: "Endgame", 
	Year: 2019, 
	Rating: 8.8, 
}}
func GetAllMoviesHandler(c echo.Context) error {
	y:= c.QueryParam("year")

	if y == "" {
	return c.JSON(200, movies)
	}

	year, err:= strconv.Atoi(y)
	if err != nil{
		return c.JSON(400, map[string]string{"message":"error type"})
	}

	findMovieByYear := []Movie{}
	for _,m:= range movies{
		if(year  == m.Year){
			findMovieByYear = append(findMovieByYear, m)
		}
	}

	if len(findMovieByYear) == 0{
	return c.JSON(404, map[string]string{"message": "not found movie"})
	}

	return c.JSON(200,findMovieByYear)
}
func GetMoviebyId(c echo.Context) error {
	id:= c.Param("id")

	for _,m:= range movies{
		if(m.ImdbID == id){
			return c.JSON(200, m)

		}
	}
	return c.JSON(404, map[string]string{"message": "not found movie"})
}
func PostMovie(c echo.Context) error {
	m:= new(Movie)
	if err:= c.Bind(m); err != nil {
		return err
	}
	movies = append(movies, *m)
	return c.JSON(200,m)

}