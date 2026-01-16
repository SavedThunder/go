package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PATCH("albums/:id/", updateAlbumPriceByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
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
	message := fmt.Sprintf("album with id: %v not found", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
}

type albumUpdatePriceBody struct {
	Price float64 `json:"price" binding:"required"`
}

func updateAlbumPriceByID(c *gin.Context) {
	var patchAlbumPrice albumUpdatePriceBody
	if err := c.ShouldBindJSON(&patchAlbumPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	log.Println("New Price", patchAlbumPrice.Price, "for id", id)

	// NOTE: range returns two things index, and a copy of the item, not the actual item.
	// Need to pass a reference to the item albums[i] does this for us

	for i, a := range albums {
		if a.ID == id {
			albums[i].Price = patchAlbumPrice.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	message := fmt.Sprintf("albmus with id: %v not found", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
}
