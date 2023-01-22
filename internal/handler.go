package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

// GetDays writes counted days from url parameter into response body
func GetDays(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil || year == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid year entered"})
		log.Println("Invalid year entered")
		return
	}
	days := countDays(year)
	c.String(http.StatusOK, days)
}

func PageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "You can only use '/when/:year' url in this app"})
}

func countDays(year int) string {
	paramDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	formattedNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	difference := int(math.Abs(paramDate.Sub(formattedNow).Hours() / 24))
	if paramDate.Before(formattedNow) {
		return fmt.Sprintf("Days gone: %s", strconv.Itoa(difference))
	} else {
		return fmt.Sprintf("Days left: %s", strconv.Itoa(difference))
	}
}
