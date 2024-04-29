package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_cache(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get lru cache end point"})
}

func Update_cache() {

}
