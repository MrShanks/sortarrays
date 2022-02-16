package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sortarray/model"
)

type Config struct {
	User     string
	Password string
	Hostname string
	Port     string
	Database string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Hostname, config.Port, config.Database)
	return connectionString
}

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	Connector.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Array{})

	log.Println("Connection was successful!!")
	return nil
}
