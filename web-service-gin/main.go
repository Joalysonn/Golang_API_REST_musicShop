package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blackstar", Artist: "David Bowie", Price: 78.22},
	{ID: "2", Title: "Wolf", Artist: "Tyler", Price: 68.99},
	{ID: "3", Title: "Javelin", Artist: "Sufjan", Price: 22.26},
	{ID: "4", Title: "Lahai", Artist: "Sampha", Price: 18.12},
	{ID: "5", Title: "Ye", Artist: "Kanye", Price: 108.19},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
