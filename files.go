package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var validExtensions = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
}

func getImagesNames(path string) []*Picture {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("There was an error getting the files from the given path: %v\n", err)
	}

	pictures := make([]*Picture, 0)

	for _, file := range files {
		extension := filepath.Ext(file.Name())
		filename := getFilename(file.Name(), extension)
		if isPicture(extension) {
			picture := NewPicture(filename, extension, path)
			pictures = append(pictures, picture)
		}
	}

	return pictures
}

func getFilename(filename, extension string) string {
	return strings.TrimSuffix(filepath.Base(filename), extension)
}

func isPicture(extension string) bool {
	_, ok := validExtensions[strings.ToLower(extension)]
	return ok
}
