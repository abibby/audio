package audio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/faiface/beep"
)

const SampleRate beep.SampleRate = 44100

var stdinBuffer *bufio.Reader

func Write(value float64) error {
	_, err := fmt.Printf("%f\n", value)
	return err
}

func Read() (float64, error) {
	if stdinBuffer == nil {
		stdinBuffer = bufio.NewReader(os.Stdin)
	}
	s, err := stdinBuffer.ReadString('\n')
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
	if len(os.Args)-1 < index {
		return 0, fmt.Errorf("no arg at %d", index)
	}
	value, err := strconv.ParseFloat(os.Args[index+1], 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}