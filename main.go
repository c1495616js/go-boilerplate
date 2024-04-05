package main

import (
	"gin-boilerplate/config"
	"gin-boilerplate/infra/database"
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/migrations"
	"gin-boilerplate/routers"
	"time"

	"github.com/spf13/viper"
)

func main() {

	//set timezone
	// default PST
	viper.SetDefault("SERVER_TIMEZONE", "America/Los_Angeles")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN := config.DbConfiguration()

	if err := database.DbConnection(masterDSN, replicaDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	//later separate migration
	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
