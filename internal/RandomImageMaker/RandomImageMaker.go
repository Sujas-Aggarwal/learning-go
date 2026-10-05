package randomimagemaker

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"strconv"
)

type Color struct {
	Red   byte
	Green byte
	Blue  byte
}

func createRandomColor() Color {
	return Color{
		byte(rand.IntN(256)),
		byte(rand.IntN(256)),
		byte(rand.IntN(256)),
	}
}

func MakeRandomImage() {
	profile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer profile.Close()

	if err := pprof.StartCPUProfile(profile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	args := os.Args

	var scale int32 = 10

	if len(args) > 1 {
		value, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		scale = int32(value)
	}

	width := int(10 * scale)
	height := int(10 * scale)

	fileName := "image.ppm"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// P6 header
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)

	if _, err := file.WriteString(header); err != nil {
		log.Fatal(err)
	}

	// Pixel data
	buffer := make([]byte, 0, width*3)

	for range height {
		buffer = buffer[:0]

		for range width {
			color := createRandomColor()

			buffer = append(
				buffer,
				color.Red,
				color.Green,
				color.Blue,
			)
		}

		if _, err := file.Write(buffer); err != nil {
			log.Fatal(err)
		}
	}
}
