package main

import (
	"github.com/Omelman/task-management/api/config"
	"github.com/Omelman/task-management/api/handlers"
	"github.com/Omelman/task-management/api/repo/postgres"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

const defaultConfigPath = "./../config.json"

func main() {
	err := config.Load(defaultConfigPath)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to laod logger: %s", err.Error())
	}

	err = postgres.Load(config.Get().Postgres, logger.Sugar())
	if err != nil {
		logger.Fatal("Failed to connect to postgres", zap.Error(err))
	}
	/*
		err = repo.Load()
		if err != nil {
			logger.Fatal("Failed to initialize postgres repo", zap.Error(err))
		}
	*/
	server := &http.Server{
		Addr:    config.Get().ListenURL,
		Handler: handlers.NewRouter(),
	}

	logger.Info("Listening...", zap.String("listen_url", config.Get().ListenURL))
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Failed to initialize HTTP server", zap.Error(err))
		os.Exit(1)
	}
}
