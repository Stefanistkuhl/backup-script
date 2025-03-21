package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type UploadWatcher struct {
	config     Config
	done       chan struct{}
	isShutdown bool
}

func NewUploadWatcher(config Config) *UploadWatcher {
	return &UploadWatcher{
		config: config,
		done:   make(chan struct{}),
	}
}

func (w *UploadWatcher) Start() {
	var lastChange time.Time
	var previousFiles map[string]time.Time
	firstRun := true
	hadChanges := false

	checkInterval, err := time.ParseDuration(w.config.CheckIntervall)
	if err != nil {
		fmt.Printf("Error parsing check interval '%s': %v. Using default 5 seconds.\n", w.config.CheckIntervall, err)
		checkInterval = 5 * time.Second
	}

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-w.done:
			return
		case <-ticker.C:
			currentFiles := make(map[string]time.Time)
			hasChanged := false

			err := filepath.Walk(w.config.UploadDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() { // Only track files, not directories
					currentFiles[path] = info.ModTime()
				}
				return nil
			})

			if err != nil {
				fmt.Printf("Error walking directory: %v\n", err)
				continue
			}

			if firstRun {
				previousFiles = currentFiles
				lastChange = time.Now()
				firstRun = false
				continue
			}

			// Compare current files with previous files
			for path, modTime := range currentFiles {
				prevModTime, exists := previousFiles[path]
				if !exists || modTime.After(prevModTime) {
					hasChanged = true
					hadChanges = true
					break
				}
			}

			// Check for deleted files
			for path := range previousFiles {
				if _, exists := currentFiles[path]; !exists {
					hasChanged = true
					hadChanges = true
					break
				}
			}

			if hasChanged {
				lastChange = time.Now()
			} else if hadChanges && !lastChange.IsZero() && time.Since(lastChange) >= checkInterval {
				w.handleUploadComplete()
				hadChanges = false // Reset for next batch of changes
			}

			previousFiles = currentFiles
		}
	}
}

func (w *UploadWatcher) Shutdown() {
	if !w.isShutdown {
		w.isShutdown = true
		close(w.done)
	}
}

func (w *UploadWatcher) handleUploadComplete() {
	// TODO: Implement upload completion handling
	fmt.Println("Upload complete! Handle the uploaded files here")
}
