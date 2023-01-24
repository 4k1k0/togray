package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

func process(wg *sync.WaitGroup, picture *Picture) {
	log.Printf("Processing %s%s\n", picture.FileName, picture.Extension)

	if _, err := os.Stat(picture.PathNew); os.IsNotExist(err) {
		os.MkdirAll(picture.PathNew, os.ModePerm)
	}

	toGray(picture)
	toMono(picture)
	toSeparate(picture)
	wg.Done()
}

func toGray(picture *Picture) {
	log.Printf("Processing %s%s: Gray\n", picture.FileName, picture.Extension)
	convertType := "gray"
	input, output := picture.GetInputAndOutputNames(convertType)
	cmd := exec.Command("convert", input, "-colorspace", "Gray", output)
	toImage(cmd, picture, convertType)
}

func toMono(picture *Picture) {
	log.Printf("Processing %s%s: Monochrome\n", picture.FileName, picture.Extension)
	convertType := "mono"
	input, output := picture.GetInputAndOutputNames(convertType)
	cmd := exec.Command("convert", input, "-monochrome", output)
	toImage(cmd, picture, convertType)
}

func toSeparate(picture *Picture) {
	log.Printf("Processing %s%s: Separate\n", picture.FileName, picture.Extension)
	convertType := "separate"
	input, output := picture.GetInputAndOutputNames(convertType)
	cmd := exec.Command("convert", input, "-separate", output)
	toImage(cmd, picture, convertType)
}

func toImage(cmd *exec.Cmd, picture *Picture, convertType string) {
	_, err := cmd.Output()

	if err != nil {
		log.Fatalf("There was an error converting %q to %s: %v\n", picture.FileName, convertType, err)
	}

	log.Printf("Done: %q\n", picture.FileName)
}
