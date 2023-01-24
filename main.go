package main

import (
	"flag"
)

func main() {
	path := flag.String("path", ".", "the input path")
	flag.Parse()
	Run(*path)
}
