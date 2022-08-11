package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// album represents the data about the record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json: "price"`
}

var albums = []album{
    // sample data
    {
        ID: "1",
        Title: "After Dark",
        Artist: "Mr. Kitty",
        Price: 56.99,
    },
    {
        ID: "2",
        Title: "Where is my mind?",
        Artist: "Pixies",
        Price: 59.99,
    },
    {
        ID: "3",
        Title: "Sarah Vaughan and Clifford Brown",
        Artist: "Sarah Vaughan",
        Price: 48.99,
    },
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    // run on localhost:8080
    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    // Call BindJSON to bind the recieved JSON in
    // newAlbum
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    // Add the newAlbum to the slice
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    // getAlbumsByID returns an album by the id which
    // matches the ID provided.
    id := c.Param("id")
    // Loop over the albums to find the specific ID
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
