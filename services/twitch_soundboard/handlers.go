package twitch_soundboard

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

func (s *Service) OnTwitchPrivateMessage(message twitch.PrivateMessage) {
	cfg, err := s.Config()
	if err != nil {
		s.log.Error("getting config", "error", err)
		return
	}

	group := cfg.Group(keyOnTwitchPrivateMessage)

	for _, trigger := range group.Keys() {
		// does the message contain the trigger?
		if !strings.Contains(message.Message, trigger) {
			continue
		}

		sounds := group.GetStrings(trigger)
		soundPath, _ := pickRandomString(sounds)

		go func() {
			if err = s.playAudioFile(soundPath); err != nil {
				s.log.Error("playing sound file", "error", err)
			}
		}()

		s.log.Info().
			Str("chat message", message.Message).
			Str("trigger", trigger).
			Str("sound", soundPath).
			Msg("playing audio")

		if s.notification != nil {
			s.notification.Notify("Twitch", fmt.Sprintf("playing %v", soundPath), "/home/gravestench/Downloads/twitch_favicon.png")
		}

		break
	}
}

func pickRandomString(slice []string) (string, error) {
	// Check if the slice is empty
	if len(slice) == 0 {
		return "", fmt.Errorf("the provided slice is empty")
	}

	// Generate a random index
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(slice))

	// Return the random string
	return slice[randomIndex], nil
}
