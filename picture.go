package main

import (
	"fmt"
)

type Picture struct {
	FileName  string
	Extension string
	NewName   string
	Path      string
	PathNew   string
}

func NewPicture(filename, extension, path string) *Picture {
	return &Picture{
		FileName:  filename,
		Extension: extension,
		Path:      path,
		PathNew:   fmt.Sprintf("%s/results", path),
	}
}

func (p *Picture) GetInputAndOutputNames(convertType string) (input, output string) {
	input = fmt.Sprintf("%s/%s%s", p.Path, p.FileName, p.Extension)
	output = fmt.Sprintf("%s/%s_%s%s", p.PathNew, p.FileName, convertType, p.Extension)
	return
}
