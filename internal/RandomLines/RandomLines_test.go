package randomlines

import "testing"

func TestCreateRandomPath(t *testing.T) {
	for range 100000 {
		randomLine := CreateRandomPath()
		if randomLine.X < 0 || randomLine.Y < 0 || randomLine.X >= 10 || randomLine.Y >= 10 {
			t.Error("Some Error Occured with Create Random Path Function.")
		}
	}
}
