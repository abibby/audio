package main

import (
	"io"
	"log"

	"github.com/abibby/audio"
)

func main() {
	f, err := audio.FloatArg(0)
	if err != nil {
		log.Fatal(err)
	}
	for {
		value, err := audio.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		audio.Write(value * f)
	}
}
