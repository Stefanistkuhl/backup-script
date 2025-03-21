package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func checkAndMakeDirs(config Config) {
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
		os.MkdirAll(config.UploadDir, os.ModePerm)
	}
	if !generationsDirExists {
		fmt.Println(config.GenerationsDir, " was not found and therefore created")
		os.MkdirAll(config.GenerationsDir, os.ModePerm)
	}
	checkAndGenerateGenerationDirs(config)

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

func checkAndGenerateGenerationDirs(config Config) {
	for i := range config.GenerationCount {
		genDirPath := filepath.Join(config.GenerationsDir, config.GenerationsDirNamePrefix+strconv.Itoa(i))
		genDirExists, err := dirExists(genDirPath)
		if err != nil {
			panic(err)
		}
		if !genDirExists {
			fmt.Println(genDirPath, " was not found and therefore created")
			os.MkdirAll(genDirPath, os.ModePerm)
		}

	}
}
