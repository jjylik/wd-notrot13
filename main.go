package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"jjylik/wd-notrot13/affine"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func getBytes(img *image.Gray) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := png.Encode(buff, img)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func getTextFromImage(img *image.Gray) (string, error) {
	imageBytes, err := getBytes(img)
	if err != nil {
		return "", err
	}
	client := gosseract.NewClient()
	client.SetImageFromBytes(imageBytes)
	text, err := client.Text()
	client.Close()
	return text, err
}

func createGrayscaleHighContrastCopyImage(source image.Image) *image.Gray {
	bounds := source.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	blackAndWhiteImage := image.NewGray(bounds)
	backgroundColor := color.RGBA{252, 245, 229, 255}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			colorFromParchment := source.At(x, y)
			if colorFromParchment == backgroundColor {
				blackAndWhiteImage.Set(x, y, color.White)
			} else {
				// The other color is the text color
				blackAndWhiteImage.Set(x, y, color.Black)
			}
		}
	}
	return blackAndWhiteImage
}

func main() {
	infile, err := os.Open("./parchment.png")
	if err != nil {
		panic(err)
	}
	defer infile.Close()
	originalPng, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}
	grayscaleCopy := createGrayscaleHighContrastCopyImage(originalPng)
	text, err := getTextFromImage(grayscaleCopy)
	if err != nil {
		panic(err)
	}
	decryptedText, err := affine.Decrypt(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(decryptedText)
}
