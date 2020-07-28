package cities

import (
	"aviasales/src/settings"
	"aviasales/src/stores"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var ticker *time.Ticker
var updaterTaskDone chan bool

func InitPeriodicalTask() {
	ticker = time.NewTicker(settings.GetSetting().PeriodicalTaskTime)
	updaterTaskDone = make(chan bool)
}

func RunTasks() {
	go updaterTask()
}

// Переодическая задача
func updaterTask() {
	for {
		select {
		case <-updaterTaskDone:
			ticker.Stop()
			return
		case currentTime := <-ticker.C:
			fmt.Println("UpdaterTask at", currentTime)
			getDataFromTravelPayouts()
		}
	}
}

// Получение данных с suggest.travelpayouts.com
func getDataFromTravelPayouts() {
	response, err := http.Get(settings.GetSetting().ApiUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result map[string]int
	jsonErr := json.NewDecoder(response.Body).Decode(&result)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}
	stores.GetDataBase().IATA = result
}
