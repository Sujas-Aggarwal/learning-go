package main

import randomimagemaker "first_project/internal/RandomImageMaker"

func main() {
	randomimagemaker.MakeRandomImage(
		5000,
		8,
		"image.ppm",
	)
}
