package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wgir/gapsi-todo/internal/application"
	"github.com/wgir/gapsi-todo/internal/infrastructure/config"
	firestoreRepo "github.com/wgir/gapsi-todo/internal/infrastructure/db/firestore"
	"github.com/wgir/gapsi-todo/internal/infrastructure/logger"
	"github.com/wgir/gapsi-todo/internal/infrastructure/web"
	"github.com/wgir/gapsi-todo/internal/infrastructure/web/handler"
	"go.uber.org/zap"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize Logger
	log, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	// 3. Initialize Firestore Client
	// If emulator host is set in config, put it in the system env so the library sees it
	if cfg.FirestoreEmulatorHost != "" {
		os.Setenv("FIRESTORE_EMULATOR_HOST", cfg.FirestoreEmulatorHost)
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatal("Failed to create firestore client", zap.Error(err))
	}
	defer client.Close()

	// 4. Dependency Injection
	repo := firestoreRepo.NewTaskRepository(client)
	service := application.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(service, log)
	router := web.NewRouter(taskHandler)

	// 5. Start HTTP Server
	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	// Channel for errors and signals
	serverErrors := make(chan error, 1)
	go func() {
		log.Info("Server is starting", zap.String("port", cfg.AppPort))
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// 6. Graceful Shutdown
	select {
	case err := <-serverErrors:
		log.Fatal("Server error", zap.Error(err))

	case sig := <-shutdown:
		log.Info("Starting shutdown", zap.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Error("Could not stop server gracefully", zap.Error(err))
			if err := server.Close(); err != nil {
				log.Fatal("Could not stop server", zap.Error(err))
			}
		}
		log.Info("Server stopped gracefully")
	}
}
