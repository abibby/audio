package main

import (
	"github.com/abibby/audio"
)

func main() {
	f, err := audio.FloatArg(0)
	if err != nil {
		f = 0.5
	}
	audio.Transform(func(value float64) float64 {
		if value > f {
			return 1
		}
		return 0
	})
}
