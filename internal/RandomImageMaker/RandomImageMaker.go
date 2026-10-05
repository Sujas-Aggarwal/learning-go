package randomimagemaker

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

func WriteLine(file *os.File, line ...any) {
	if _, err := file.WriteString(fmt.Sprintln(line...)); err != nil { // we will always go the next line, no other choice
		log.Fatal(err)
	}
}

func createRandomColor() Color {
	return Color{rand.IntN(256), rand.IntN(256), rand.IntN(256)}
}

func MakeRandomImage() {
	args := os.Args
	var SCALE int32 = 10
	if len(args) > 1 {
		scale, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		SCALE = int32(scale)
	}
	var WIDTH int32 = 10 * SCALE
	var HEIGHT int32 = 10 * SCALE
	var FILE_NAME string = "image.ppm"
	var FORMAT string = "P3"
	var MAX_COLOR_VALUE string = "255"

	// truncating if the file exists
	_, er := os.Stat(FILE_NAME)
	if er == nil {
		err := os.Truncate(FILE_NAME, 0)
		if err != nil {
			log.Fatalf("Failed to clear file: %v", err)
		}
	}
	file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed when the function finishes

	// Append your text to the file
	WriteLine(file, FORMAT)
	WriteLine(file, WIDTH, HEIGHT)
	WriteLine(file, MAX_COLOR_VALUE)
	// so far we have basically added the details, now we just need to add the values
	// we can basically loop now and add pixels easily one after another
	for range WIDTH {
		valuesToPush := []int{}
		for range HEIGHT {
			color := createRandomColor()
			// so whar we will be doing here is, we will first save the entire line to write in memory and then dump it in the file to minimse IO operations.
			valuesToPush = append(valuesToPush, []int{color.Red, color.Green, color.Blue}...)
		}
		WriteLine(file, valuesToPush) // appending at once
	}
}
