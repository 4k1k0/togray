package main

import (
	"log"
	"runtime"
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

	maxProcess := getMaxProcess()
	waitChan := make(chan struct{}, maxProcess)
	wg := sync.WaitGroup{}
	wg.Add(numberOfPictures)

	for _, picture := range pictures {
		waitChan <- struct{}{}
		go func(picture *Picture) {
			process(&wg, picture)
			<-waitChan
		}(picture)
	}

	wg.Wait()

}

func getMaxProcess() int {
	if n := runtime.NumCPU(); n > 1 {
		return n / 2
	}
	return 1
}
