package main

import (
	routers "lru_cache_backend/router"
	"lru_cache_backend/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ws", websocket.HandleWebSocket)

	lru_group := router.Group("/api/lru_cache")
	routers.Lru_Routes(lru_group)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a test end point"})
	})

	router.Run(":8000")
}
