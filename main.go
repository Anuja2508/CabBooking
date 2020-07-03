package main

import (
	"GoCab/model"
	"GoCab/server"
	"fmt"

	"github.com/labstack/gommon/log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/spf13/viper"
)

func main() {

	log.Print("GET CONFIG")

	config.AddConfigPath(".")
	config.SetConfigName("config")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error getting config from file: %s", err))
	}
	log.Print("CONNECT TO DATABASE")
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.GetString("postgres.host"),
			config.GetString("postgres.port"),
			config.GetString("postgres.user"),
			config.GetString("postgres.dbname"),
			config.GetString("postgres.password"),
			config.GetString("postgres.sslmode"),
		),
	)
	if err != nil {
		log.Panic("COULDN'T CONNECT TO DATABASE " + err.Error())
	}

	db.AutoMigrate(&model.User{},
		&model.Cab{},
		&model.Booking{})

	apiserver, err := server.New(db)

	if err != nil {
		panic(fmt.Errorf("fatal error setting up server: %s", err))
	}

	port := ":" + config.GetString("backend.port")
	apiserver.Logger.Fatal(apiserver.Start(port))
}
