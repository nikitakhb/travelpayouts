package cities

import (
	"aviasales/src/settings"
	"aviasales/src/stores"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var ticker *time.Ticker
var updaterTaskDone chan bool

func InitPeriodicalTask() {
	ticker = time.NewTicker(settings.GetSetting().PeriodicalTaskTime)
	updaterTaskDone = make(chan bool)
}

func RunTasks() {
	go updaterTask()
	getDataFromTravelPayouts()
}

func StopTasks() {
	updaterTaskDone <- false
}

// Переодическая задача (конечно тут нет мьютекса какого нибудь на случай если задача
// будет долго исполняться, и запуститься еще одна задача =)) и они будут конкурентно писать в стор.)
func updaterTask() {
	for {
		select {
		case <-updaterTaskDone:
			ticker.Stop()
			return
		case <-ticker.C:
			getDataFromTravelPayouts()
		}
	}
}

// Получение данных с suggest.travelpayouts.com
func getDataFromTravelPayouts() {
	logRecord := log.WithFields(log.Fields{
		"package":  "cities",
		"func":     "getDataFromTravelPayouts",
		"endpoint": settings.GetSetting().ApiUrl,
	})

	logRecord.WithFields(log.Fields{
		"status": "run",
	}).Info("Запущена задача получения данных с эндпоинта")

	response, err := http.Get(settings.GetSetting().ApiUrl)
	if err != nil {
		logRecord.WithFields(log.Fields{
			"status": "error",
			"error":  err.Error(),
		}).Error("Произошла ошибка при получении данных с эндпоинта")
		return
	}
	var result map[string]int
	jsonErr := json.NewDecoder(response.Body).Decode(&result)
	if jsonErr != nil {
		logRecord.WithFields(log.Fields{
			"status":        "error",
			"error":         jsonErr.Error(),
			"response_body": response.Body,
		}).Error("Произошла ошибка при преобразовании данных")
		return
	}
	stores.GetDataBase().IATA = result
	log.WithFields(log.Fields{
		"package":  "cities",
		"func":     "getDataFromTravelPayouts",
		"status":   "completed",
		"endpoint": settings.GetSetting().ApiUrl,
	}).Info("Успешно получены данные с эндпоинта.", settings.GetSetting().ApiUrl)
}
