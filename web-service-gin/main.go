package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PATCH("albums/:id/price", updateAlbumPriceByID)
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

func updateAlbumPrice(album *album, newPrice string) (a album, err error) {
	log.Println(newPrice)
	f, err := strconv.ParseFloat(newPrice, 64)
	if err != nil {
		return *album, err
	}
	album.Price = f
	return *album, err
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

func updateAlbumPriceByID(c *gin.Context) {
	id := c.Param("id")
	newPrice := c.Param("price")
	log.Println(newPrice)
	for _, a := range albums {
		if a.ID == id {
			_, err := updateAlbumPrice(&a, newPrice)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Err": err})
				return
			}

			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	message := fmt.Sprintf("albmus with id: %v not found", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
}
