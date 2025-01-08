package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Define the database conn configuration
type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}
)

var err error

// Connect to mysql with the input configuration
func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
		"utf8mb4",
		"True",
		"Local",
	)

	dbConn, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:           "mysql",
		DisableWithReturning: true,
		DSN:                  dsn,
	}), &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
		Logger:                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
