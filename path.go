package main

import (
	"log"
	"os"
)

func getBasePath(flagPath string) string {
	if flagPath != "." {
		return flagPath
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("there was an error getting current path: %v\n", err)
	}
	return path
}
