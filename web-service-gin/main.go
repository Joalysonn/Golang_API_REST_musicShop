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

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"messege": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
