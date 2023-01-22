package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zanovru/days-counting-test-task/internal"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(internal.PingPongCheckerMiddleware())
	router.GET("/when/:year", internal.GetDays)
	router.NoRoute(internal.PageNotFound)
	log.Fatal(router.Run("localhost:8080"))
}
