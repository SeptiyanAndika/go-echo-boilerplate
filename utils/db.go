package utils

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Db *gorm.DB
}

var instance *DB
var onceDb sync.Once

func GetInstanceDB() *DB {
	onceDb.Do(func() {
		mysqlConfig := Config.Database.User + ":" + Config.Database.Password + "@tcp(" + Config.Database.Server + ":" + Config.Database.Port + ")/" + Config.Database.Database
		dbConnection, _ := gorm.Open("mysql", mysqlConfig+"?charset=utf8&parseTime=True&loc=Local")
		instance = &DB{Db: dbConnection}
		//defer dbConnection.Close()
	})
	return instance
}
