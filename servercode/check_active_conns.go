package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration with error %v", err)
	}
	fmt.Println(config)
	checAndMakeDirs(config)
	//lsof -i
	//https://stackoverflow.com/questions/41259191/golang-periodically-checking-open-tcp-connections-at-a-port#41261128
}
