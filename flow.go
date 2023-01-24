package main

import (
	"log"
	"sync"
)

func Run(flagPath string) {
	path := getBasePath(flagPath)
	pictures := getImagesNames(path)
	numberOfPictures := len(pictures)
	if numberOfPictures == 0 {
		log.Printf("There are no images to process")
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(numberOfPictures)

	for _, picture := range pictures {
		go func(picture *Picture) {
			process(&wg, picture)
		}(picture)
	}

	wg.Wait()

}
