package main

import (
	"log"
	"os"

	"github.com/etowett/datsimple/backend/db"
	"github.com/etowett/datsimple/backend/logger"
	"github.com/etowett/datsimple/backend/repos"
	"github.com/etowett/datsimple/backend/services"
	appRouter "github.com/etowett/datsimple/backend/web"
)

func main() {
	if os.Getenv("LOG_FILE") != "" {
		logFile, err := os.OpenFile(os.Getenv("LOG_FILE"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			logger.Fatalf("log file open error: %v", err)
		}

		defer logFile.Close()

		logger.SetOutput(logFile)
	}

	dbManager := db.NewPostgresDBManager()
	defer dbManager.Close()

	err := dbManager.Ping()
	if err != nil {
		log.Fatalf("error db ping: %v", err)
	}

	costService := services.NewDailyCostService(repos.NewDailyCostRepository())

	appRouter.BuildRouter(
		dbManager,
		costService,
	).Run()
}
