package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// / GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PutAlbums(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")

	// Define a struct to represent the partial update
	type PartialUpdate struct {
		Title  *string  `json:"title"`
		Artist *string  `json:"artist"`
		Price  *float64 `json:"price"`
	}

	// Parse the partial update data from the request body
	var partialUpdate PartialUpdate
	if err := c.BindJSON(&partialUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the album with the matching ID
	var found bool
	for i := range albums {
		if albums[i].ID == id {
			found = true
			// Apply the partial update to the album
			if partialUpdate.Title != nil {
				albums[i].Title = *partialUpdate.Title
			}
			if partialUpdate.Artist != nil {
				albums[i].Artist = *partialUpdate.Artist
			}
			if partialUpdate.Price != nil {
				albums[i].Price = *partialUpdate.Price
			}
			break
		}
	}

	// Check if the album with the given ID was found
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully"})
}
