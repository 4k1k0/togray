package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disintegration/gift"
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

func (p *Picture) GetFileNameWithPath() string {
	return fmt.Sprintf("%s/%s%s", p.Path, p.FileName, p.Extension)
}

func (p *Picture) GenFileNameToSave(style string) string {
	return fmt.Sprintf("%s/%s_%s_%d%s", p.PathNew, p.FileName, style, time.Now().Unix(), p.Extension)
}

func (p *Picture) LoadImage() image.Image {
	f, err := os.Open(p.GetFileNameWithPath())
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func (p *Picture) ProcessImage() {
	src := p.LoadImage()

	for style, filter := range filters {
		g := gift.New(filter)
		dst := image.NewNRGBA(g.Bounds(src.Bounds()))
		g.Draw(dst, src)
		p.SaveImage(style, dst)
	}
}

func (p *Picture) SaveImage(style string, img image.Image) {
	f, err := os.Create(p.GenFileNameToSave(style))
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()

	switch strings.ToLower(p.Extension) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(f, img, nil)
	case ".png":
		err = png.Encode(f, img)

	}

	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}

}
