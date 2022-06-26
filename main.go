package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type game struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Platform string  `json:"platform"`
	Year     string  `json:"year"`
	Price    float64 `json:"price"`
}

var games = []game{
	{ID: "1", Title: "Horizon: Zero Dawn", Genre: "Action", Platform: "PS4", Year: "2017", Price: 250.00},
	{ID: "2", Title: "God of War", Genre: "Adventure", Platform: "PS4", Year: "2018", Price: 250.00},
	{ID: "3", Title: "Spiderman: Miles Morales, Ultimate Edition", Genre: "Action", Platform: "PS5", Year: "2022", Price: 700.00},
	{ID: "4", Title: "Horizon: Forbidden West", Genre: "Action", Platform: "PS5", Year: "2022", Price: 550.00},
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.POST("/games", addGame)
	router.GET("/games/:id", getGameById)

	router.Run("localhost:8080")
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func getGameById(c *gin.Context) {
	id := c.Param("id")

	for _, game := range games {
		if game.ID == id {
			c.IndentedJSON(http.StatusOK, game)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found"})
}

func addGame(c *gin.Context) {
	var newGame game
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}
