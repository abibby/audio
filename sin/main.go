package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"time"

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
		in := bufio.NewReader(os.Stdin)
		for {
			// lastHz := hz
			value, err := audio.Read(in)
			if err != nil {
				if err == io.EOF {
					return
				}
			}
			nextHz = value
			time.Sleep(time.Second / time.Duration(audio.SampleRate))
		}
	}()

	for {
		t = float64(i) / float64(sampleRate)
		// fmt.Print(i, " - ")
		value = math.Sin(t * math.Pi * hz)
		audio.Write(os.Stdout, value)
		// i = (i + 1) % (sampleRate * uint(hz))
		i++
		if lastValue > 0 && value <= 0 {
			hz = nextHz
			i = 0
		}
		lastValue = value
	}
}
