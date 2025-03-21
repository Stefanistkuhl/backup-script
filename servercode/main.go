package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := setup()

	watcher := NewUploadWatcher(config)

	// Handle graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nShutting down...")
		watcher.Shutdown()
	}()

	watcher.Start()
}

func setup() Config {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration with error %v", err)
	}
	fmt.Println(config)
	checkAndMakeDirs(config)
	return config
}
