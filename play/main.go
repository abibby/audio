package main

import (
	"log"
	"time"

	"github.com/abibby/audio"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// type SinStreamer struct {
// 	SampleRate beep.SampleRate
// 	hz         float64
// 	t          uint64
// }

// var _ beep.Streamer = &SinStreamer{}

// func (s *SinStreamer) Stream(samples [][2]float64) (n int, ok bool) {
// 	for i := range samples {
// 		t := float64(s.t) / float64(s.SampleRate)
// 		samples[i][0] = math.Sin(t * math.Pi * s.hz)
// 		samples[i][1] = math.Sin(t * math.Pi * s.hz)
// 		s.t++
// 	}
// 	return len(samples), true
// }

// func (s *SinStreamer) Err() error {
// 	return nil
// }

// func Sin(hz float64) *SinStreamer {
// 	return &SinStreamer{
// 		hz:         hz,
// 		SampleRate: beep.SampleRate(44100),
// 	}
// }

func main() {
	speaker.Init(audio.SampleRate, audio.SampleRate.N(time.Second/10))
	var value float64
	var err error
	speaker.Play(beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {

		for i := range samples {
			value, err = audio.Read()
			if err != nil {
				log.Print(err)
			}
			samples[i][0] = value
			samples[i][1] = value
		}
		return len(samples), true
	}))
	select {}
}
