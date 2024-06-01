package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ilhamgepe/mygram/cmd/api"
	"github.com/ilhamgepe/mygram/config"
	"github.com/ilhamgepe/mygram/db/postgres"
)

func init() {
	if err := config.LoadConfig("./.env"); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}

func main() {
	db := postgres.NewPostgresDB()
	api := api.NewApi(db, config.Get.Addr, config.Get.GIN_MODE)

	go func() {
		log.Fatal(api.Run())
	}()
	fmt.Println("server running on port http://localhost:" + config.Get.Addr)

	// Graceful Shutdown
	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, os.Interrupt)
	fmt.Println("signal received: ", <-stopC)

	if err := api.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown error: %s", err)
	}

	fmt.Println("server shutdown successfully")
	os.Exit(0)
}
