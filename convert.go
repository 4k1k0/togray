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

	log.Printf("Done: %q\n", picture.FileName)
	wg.Done()
}

func toGray(picture *Picture) {
	convertType := "gray"
	input, output := picture.GetInputAndOutputNames(convertType)
	cmd := exec.Command("convert", input, "-colorspace", "Gray", output)
	toImage(cmd, picture, convertType)
}

func toMono(picture *Picture) {
	convertType := "mono"
	input, output := picture.GetInputAndOutputNames(convertType)
	cmd := exec.Command("convert", input, "-monochrome", output)
	toImage(cmd, picture, convertType)
}

func toSeparate(picture *Picture) {
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
}
