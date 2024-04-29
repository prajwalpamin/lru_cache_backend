package routers

import (
	"lru_cache_backend/controller"

	"github.com/gin-gonic/gin"
)

func Lru_Routes(router *gin.RouterGroup) {
	router.GET("/get_cache", controller.Get_all_cache)
}
