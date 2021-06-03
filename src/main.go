//
// EPITECH PROJECT, 2021
// Auction-Site
// File description:
// main
//

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Auction Site",
		})
	})
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "There is the items.",
		})
	})
	router.Run()
}
