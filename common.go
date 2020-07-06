package audio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/faiface/beep"
)

const SampleRate beep.SampleRate = 44100

func Write(out io.Writer, value float64) error {
	_, err := fmt.Fprintf(out, "%f\n", value)
	return err
}

func Read(in *bufio.Reader) (float64, error) {
	s, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseFloat(s[:len(s)-1], 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func FloatArg(index int) (float64, error) {
	if len(os.Args) <= index+1 {
		return 0, fmt.Errorf("no arg at %d", index)
	}
	value, err := strconv.ParseFloat(os.Args[index+1], 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func Transform(in *bufio.Reader, out io.Writer, f func(value float64) float64) {
	for {
		value, err := Read(in)
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		Write(out, f(value))
	}
}
