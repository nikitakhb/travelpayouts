package cities

import (
	"aviasales/src/stores"
	"github.com/gin-gonic/gin"
)

func GetCountIATAView(context *gin.Context) {
	cityCode := context.Query("iata")

	if count, success := stores.GetDataBase().IATA[cityCode]; success {
		context.JSON(200, gin.H{
			"count": count,
		})
	} else {
		context.JSON(404, "Not Found!")
	}
}
