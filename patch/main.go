package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/abibby/audio"
)

func main() {
	streams := []*bufio.Reader{}

	streams = append(streams, bufio.NewReader(os.Stdin))

	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		streams = append(streams, bufio.NewReader(f))
	}

	values := make([]float64, len(streams))

	for {
		for i, stream := range streams {
			value, err := audio.Read(stream)
			if err == io.EOF {
				return
			} else if err != nil {
				log.Print(err)
			}
			values[i] = value
		}

		audio.Write(os.Stdout, avg(values))
	}
}

func avg(nums []float64) float64 {
	if len(nums) == 0 {
		panic("no numbers in average")
	}
	total := 0.0

	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}
