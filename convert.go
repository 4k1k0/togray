package main

import (
	"log"
	"os"
	"sync"
)

func process(wg *sync.WaitGroup, p *Picture) {
	log.Printf("Processing %s%s\n", p.FileName, p.Extension)

	if _, err := os.Stat(p.PathNew); os.IsNotExist(err) {
		os.MkdirAll(p.PathNew, os.ModePerm)
	}

	p.ProcessImage()

	log.Printf("Done: %q\n", p.FileName)
	wg.Done()
}
