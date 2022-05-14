package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "This One's for You", Artist: "Luke Combs", Price: 16.99, Year: 2017, Genre: "Country"},
	{ID: "2", Title: "What You See Is What You Get", Artist: "Luke Combs", Price: 17.99, Year: 2019, Genre: "Country"},
	{ID: "3", Title: "Wild Ones", Artist: "Kip Moore", Price: 19.99, Year: 2015, Genre: "Country"},
	{ID: "4", Title: "Up All Night", Artist: "Kip Moore", Price: 12.99, Year: 2012, Genre: "Country"},
	{ID: "5", Title: "Barefoot Blue Jean Night", Artist: "Jake Owen", Price: 10.99, Year: 2011, Genre: "Country"},
	{ID: "6", Title: "Traveller", Artist: "Chris Stapleton", Price: 19.99, Year: 2015, Genre: "Country"},
	{ID: "7", Title: "Sinners Like Me", Artist: "Eric Church", Price: 14.99, Year: 2006, Genre: "Country"},
	{ID: "8", Title: "Carolina", Artist: "Eric Church", Price: 11.99, Year: 2009, Genre: "Country"},
	{ID: "9", Title: "Dangerous", Artist: "Morgan Wallen", Price: 29.99, Year: 2021, Genre: "Country"},
	{ID: "10", Title: "Rollin On the River", Artist: "CCR", Price: 9.99, Year: 1988, Genre: "Rock"},
	{ID: "11", Title: "Enema of the State", Artist: "Blink 182", Price: 5.99, Year: 1999, Genre: "Rock"},
	{ID: "12", Title: "21", Artist: "Adele", Price: 15.99, Year: 2011, Genre: "Pop"},
	{ID: "13", Title: "Nine Tonight", Artist: "Bob Seger", Price: 9.99, Year: 1981, Genre: "Rock"},
	{ID: "14", Title: "(What's the Story) Morning Glory?", Artist: "Oasis", Price: 10.99, Year: 1995, Genre: "Rock"},
}

//getAlbums responds with the list of all albums as json
func getAlbums(c *gin.Context) {
	//indentedjson serializes the struct into JSON and adds it to the response, sends 200 OK status
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	//call BindJSON to bind the received json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//append the new album to the albums slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	//locates the album whose ID matches the value sent by the client and returns i
	id := c.Param("id")
	//loop over the list of albums and return the proper ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//if the ID is not in albums - return error message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")
	var byartist []album
	for _, a := range albums {
		if a.Artist == artist {
			byartist = append(byartist, a)
		}
	}
	if len(byartist) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "artist not found"})
	} else {
		c.IndentedJSON(http.StatusOK, byartist)
	}
}

func getAlbumsByYear(c *gin.Context) {
	year := c.Param("year")
	var byartist []album
	for _, a := range albums {
		x := fmt.Sprintf("%v", a.Year)
		if x == year {
			byartist = append(byartist, a)
		}
	}
	if len(byartist) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "artist not found"})
	} else {
		c.IndentedJSON(http.StatusOK, byartist)
	}
}

func main() {
	// initialize gin router
	router := gin.Default()
	// use the GET function to associate the GET HTTP method with the /albums path
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/:artist", getAlbumsByArtist)
	// associate the POST HTTP method with the post function
	router.POST("/albums", postAlbums)
	//attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}
