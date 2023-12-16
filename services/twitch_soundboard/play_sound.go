package twitch_soundboard

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func (s *Service) playAudioFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		s.log.Error("opening audio file", "error", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		s.log.Error("playing mp3 audio file", "error", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}
