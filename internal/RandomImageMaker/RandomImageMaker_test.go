package randomimagemaker

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"testing"
	"time"
)

func TestProfileRandomImage(t *testing.T) {
	scales := []int{
		1,
		10,
		100,
		500,
		1000,
	}

	workers := []int{
		1,
		2,
		4,
		8,
		16,
		32,
		64,
		100,
	}

	profileDir := "profiles"

	if err := os.MkdirAll(profileDir, 0755); err != nil {
		t.Fatal(err)
	}

	for _, scale := range scales {
		for _, workerCount := range workers {

			t.Run(
				fmt.Sprintf("scale_%d_workers_%d", scale, workerCount),
				func(t *testing.T) {

					imageFile := filepath.Join(
						"benchmark_images",
						fmt.Sprintf(
							"image_scale_%d_workers_%d.ppm",
							scale,
							workerCount,
						),
					)

					profileFile := filepath.Join(
						profileDir,
						fmt.Sprintf(
							"scale_%d_workers_%d.prof",
							scale,
							workerCount,
						),
					)

					if err := os.MkdirAll("benchmark_images", 0755); err != nil {
						t.Fatal(err)
					}

					// Start CPU profiling.
					profile, err := os.Create(profileFile)
					if err != nil {
						t.Fatal(err)
					}

					if err := pprof.StartCPUProfile(profile); err != nil {
						profile.Close()
						t.Fatal(err)
					}

					start := time.Now()

					MakeRandomImage(
						scale,
						workerCount,
						imageFile,
					)

					elapsed := time.Since(start)

					pprof.StopCPUProfile()
					profile.Close()

					// Get output file size.
					info, err := os.Stat(imageFile)
					if err != nil {
						t.Fatal(err)
					}

					t.Logf(
						"scale=%d workers=%d duration=%v file_size=%d bytes profile=%s",
						scale,
						workerCount,
						elapsed,
						info.Size(),
						profileFile,
					)

					// Delete huge image after profiling.
					if err := os.Remove(imageFile); err != nil {
						t.Logf("failed to remove image: %v", err)
					}
				},
			)
		}
	}
}
