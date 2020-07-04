package main

import (
	"io"
	"math"

	"github.com/abibby/audio"
)

func main() {
	hz := 440.0
	nextHz := hz
	i := uint(0)
	t := 0.0
	sampleRate := uint(audio.SampleRate)
	value := 0.0
	lastValue := 0.0

	go func() {
		for {
			// lastHz := hz
			value, err := audio.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
			}
			nextHz = value
		}
	}()

	for {
		t = float64(i) / float64(sampleRate)
		// fmt.Print(i, " - ")
		value = math.Sin(t * math.Pi * hz)
		audio.Write(value)
		// i = (i + 1) % (sampleRate * uint(hz))
		i++
		if lastValue > 0 && value <= 0 {
			hz = nextHz
			i = 0
		}
		lastValue = value
	}
}
