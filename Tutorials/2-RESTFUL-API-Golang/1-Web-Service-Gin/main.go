package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []Album{
	{
		ID: "1",
		Title: "Blue Train",
		Artist: "John Coltrane",
		Price: 56.99,
	},
	{
		ID: "2",
		Title: "Jeru",
		Artist: "Gerry Mulligan",
		Price: 17.99,
	},
	{
		ID: "3",
		Title: "Sarah Vaughan and Clifford Brown",
		Artist: "Saray Vaughan",
		Price: 39.99,
	},
}

func main() {
	router := gin.Default()
	// Associating the GET HTTP method and /albums path with a handler function.
	// GET ALL Albums.
	router.GET("/albums", getAlbums)

	// Associating the GET HTTP method and /albums/:id path with a handler function.
	// GET specific album with id.
	router.GET("/albums/:id", getAlbumByID)
	
	// Associating the POST HTTP method and /albums path with a handler function.
	// Create New Album
	router.POST("/albums", postAlbums)

	fmt.Println("Running server on port: 8282")
	// Attaching the router to a run server
	router.Run("localhost:8282")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request bod.
func postAlbums(ctx *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a reponse.
func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}