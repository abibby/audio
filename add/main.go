package main

import (
	"log"

	"github.com/abibby/audio"
)

func main() {
	f, err := audio.FloatArg(0)
	if err != nil {
		log.Fatal(err)
	}
	audio.Transform(func(value float64) float64 {
		return value + f
	})
}
