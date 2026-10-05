package randomimagemaker

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"sync"
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

func MakeRandomImage(scale int, workers int, fileName string) {
	width := 10 * scale
	height := 10 * scale

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND,
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

	rowsPerWorker := height / workers
	remainder := height % workers

	var wg sync.WaitGroup
	wg.Add(workers)

	for worker := 0; worker < workers; worker++ {
		// Distribute remainder rows among the first workers.
		rows := rowsPerWorker
		if worker < remainder {
			rows++
		}

		go func() {
			defer wg.Done()

			buffer := make([]byte, 0, width*3)

			for range rows {
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
					log.Println("write error:", err)
					return
				}
			}
		}()
	}

	wg.Wait()
}
