package main

import randomimagemaker "first_project/internal/RandomImageMaker"

func main() {
	randomimagemaker.MakeRandomImage(
		500,
		1,
		"image.ppm",
	)
}
