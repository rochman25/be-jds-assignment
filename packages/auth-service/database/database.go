package database

import (
	config "auth-service/pkg"
	"gorm.io/gorm"
	"sync"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	// Create database configuration information
	conf := dbConfig{
		User: config.MysqlUser(),
		Pass: config.MysqlPass(),
		Host: config.MysqlHost(),
		Port: config.MysqlPort(),
		Name: config.MysqlDbName(),
	}

	mysql := mysqlConfig{dbConfig: conf}
	// Create only one mysql Connection, not the same as mysql TCP connection
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	// Check db connection, if exist return the memory address of the db connection
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
