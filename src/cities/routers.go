package cities

import (
	"github.com/gin-gonic/gin"
)

func CityRegister(router *gin.RouterGroup) {
	router.GET("/get_count", GetCountIATAView)
}
