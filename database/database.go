package database

import (
  "github.com/jinzhu/gorm"
)

var databaseHandle *gorm.DB

type Connector interface {
	Connect()
	Destroy()
}

func GetDatabaseHandle() *gorm.DB {
	return databaseHandle
}

func NewDatabase(mType, username, password, dbName string) MysqlDataBase {
	return MysqlDataBase {
		Type: mType,
		Username: username,
		Password: password,
		DBName: dbName,
	}
}
