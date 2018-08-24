package gfx

import (
	"image"
	"os"
)

//LoadImageFromFile Loads an image from file
func LoadImageFromFile(file string) (image.Image, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}
