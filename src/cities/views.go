package cities

import (
	"aviasales/src/stores"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCountIATAView(context *gin.Context) {
	log.WithFields(log.Fields{
		"package": "cities",
		"view":    "GetCountIATAView",
		"request": map[string]string{
			"method": context.Request.Method,
			"query":  context.Request.URL.RawQuery,
		},
	}).Info("Request")
	cityCode := context.Query("iata")

	if count, success := stores.GetDataBase().IATA[cityCode]; success {
		context.JSON(200, gin.H{
			"count": count,
		})
	} else {
		context.JSON(404, "Not Found!")
	}
}
