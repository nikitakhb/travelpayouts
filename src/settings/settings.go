package settings

import (
	"os"
	"strconv"
	"sync"
	"time"
)

type Settings struct {
	ApiUrl             string
	PeriodicalTaskTime time.Duration
}

var instanceSetting *Settings
var once sync.Once

func GetSetting() *Settings {
	once.Do(func() {
		taskUpdatePeriod, err := strconv.Atoi(getEnv("TASK_UPDATE_PERIOD", "10"))
		if err != nil {
			taskUpdatePeriod = 10
		}
		instanceSetting = &Settings{
			ApiUrl: getEnv(
				"API_URL",
				"https://suggest.travelpayouts.com/data_api?service=random_city_sample&no-cache=true"),
			PeriodicalTaskTime: time.Duration(taskUpdatePeriod) * time.Second,
		}
	})
	return instanceSetting
}

func getEnv(key, defaultValue string) string {
	if value, success := os.LookupEnv(key); success {
		return value
	}
	return defaultValue
}
