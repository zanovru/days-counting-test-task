package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func getDays(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		fmt.Println(err)
	}
	days := countDays(year)
	c.IndentedJSON(http.StatusOK, days)
}

func countDays(year int) gin.H {
	paramDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	formattedNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	difference := paramDate.Sub(formattedNow).Hours() / 24
	if paramDate.Before(formattedNow) {
		return gin.H{"Days gone": -1 * difference}
	} else {
		return gin.H{"Days left": difference}
	}
}

func main() {
	router := gin.Default()
	router.GET("/when/:year", getDays)
	router.Run("localhost:8080")
}
