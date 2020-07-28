package stores

import "sync"

// Псевдо БД
type DB struct {
	IATA map[string]int
}

var instanceDB *DB
var once sync.Once

func GetDataBase() *DB {
	once.Do(func() {
		instanceDB = &DB{}
	})
	return instanceDB
}
