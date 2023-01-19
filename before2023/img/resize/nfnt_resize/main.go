package main

import (
	"log"
	"os"
	"image/jpeg"
	"github.com/nfnt/resize"
)

func main(){
	// open "test.jpg"
	file, err := os.Open("ai.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(300, 0, img, resize.Lanczos3)

	out, err := os.Create("ai_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
