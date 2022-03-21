package main

import (
	"os"
	"sortarray/database"
	"sortarray/service"
	"sortarray/utils"
)

func main() {

	counter := 0.0

	config :=
		database.Config{
			Hostname: os.Getenv("DB_HOSTNAME"),
			Port:     "3306",
			User:     "root",
			Password: os.Getenv("DB_PASSWORD"),
			Database: "array",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	for err != nil {
		utils.RetryBackOff(counter)
		err = database.Connect(connectionString)
		counter++
	}

	service.RestController()
}
