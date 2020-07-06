package main

import (
	"bufio"
	"log"
	"os"

	"github.com/abibby/audio"
)

func main() {
	f, err := audio.FloatArg(0)
	if err != nil {
		log.Fatal(err)
	}
	audio.Transform(
		bufio.NewReader(os.Stdin),
		os.Stdout,
		func(value float64) float64 {
			return value + f
		},
	)
}
