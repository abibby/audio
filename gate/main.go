package main

import (
	"bufio"
	"os"

	"github.com/abibby/audio"
)

func main() {
	f, err := audio.FloatArg(0)
	if err != nil {
		f = 0.5
	}
	audio.Transform(bufio.NewReader(os.Stdin), os.Stdout, func(value float64) float64 {
		if value > f {
			return 1
		}
		return 0
	})
}
