package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig(file string) (Config, error) {
	var config Config

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func checAndMakeDirs(config Config) {
	uploadDirExists, err := dirExists(config.UploadDir)
	if err != nil {
		panic(err)
	}
	generationsDirExists, err := dirExists(config.GenerationsDir)
	if err != nil {
		panic(err)
	}
	if !uploadDirExists {
		fmt.Println(config.GenerationsDir, " was not found and therefore created")
		//give this proper perms
		os.MkdirAll(config.UploadDir, 770)
	}
	if !generationsDirExists {
		fmt.Println(config.GenerationsDir, " was not found and therefore created")
		//give this proper perms
		os.MkdirAll(config.GenerationsDir, 770)
	}

}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
