package main

import (
	"log"
	"os"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create a new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Archivo modificado:", event.Name)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("Archivo creado:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("Archivo eliminado:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path to the watcher.
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "."
	}
	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan bool)
}
