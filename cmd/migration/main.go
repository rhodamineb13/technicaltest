package migration

import (
	"technicaltest/config"
	"technicaltest/database"
	"technicaltest/database/entities"
)

func main() {
	config.NewConfig()
	database.ConnectDB()
	entities := []any{&entities.MyClient{}}

	if err := database.DB.AutoMigrate(entities); err != nil {
		panic(err)
	}

}
