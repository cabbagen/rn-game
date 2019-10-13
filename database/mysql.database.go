package database

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDataBase struct {
	Type         string
	Username     string
	Password     string
	DBName       string
}

func (db MysqlDataBase) Connect() {
	connectString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", db.Username, db.Password, db.DBName)
	
	if db, error := gorm.Open(db.Type, connectString); error != nil {
		log.Fatalln(error)
	} else {
		databaseHandle = db
		gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
			return "cb_" + defaultTableName;
		}
	}
}

func (db MysqlDataBase) Destroy() {
	if databaseHandle != nil {
		databaseHandle.Close()
	}
}

var DefaultMysqlConfig map[string]string = map[string]string {
	"username": "root",
	"password": "artART5201314??",
	"dbName": "cb_game",
} 
