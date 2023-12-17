package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gravestench/servicemesh"
	"github.com/hegedustibor/htgo-tts/voices"

	"github.com/gravestench/servicemesh-examples/services/text_to_speech"
)

// this service will just connect the TTS to the twitch integration service
type glueService struct {
	tts                 text_to_speech.ConvertsTextToSpeech
	lastPersonThatSpoke string
	startupTime         time.Time
	onJoinDelay         time.Duration // prevent onJoin messages for a duration
}

func (g *glueService) OnTwitchPrivateMessage(message twitch.PrivateMessage) {
	if strings.Contains(message.Message, "http") {
		return
	}

	if g.lastPersonThatSpoke != message.User.Name {
		g.lastPersonThatSpoke = message.User.Name

		name := strings.ReplaceAll(message.User.Name, "_", "")
		g.tts.SetVoice(voices.EnglishUK)
		g.tts.Speak(name + " says: ")
	}

	g.tts.SetVoice(voices.EnglishAU)
	g.tts.Speak(message.Message)
}

func (g *glueService) OnTwitchUserJoinMessage(message twitch.UserJoinMessage) {
	if time.Since(g.startupTime) < g.onJoinDelay {
		return
	}

	g.tts.SetVoice(voices.EnglishAU)
	g.tts.Speak(fmt.Sprintf("user %s has joined the chat", message.User))
}

func (g *glueService) DependenciesResolved() bool {
	return g.tts != nil
}

func (g *glueService) ResolveDependencies(mesh servicemesh.Mesh) {
	for _, service := range mesh.Services() {
		if candidate, ok := service.(text_to_speech.ConvertsTextToSpeech); ok {
			g.tts = candidate
		}
	}
}

func (g *glueService) Init(mesh servicemesh.Mesh) {
	// do nothing
}

func (g *glueService) Name() string {
	return "glue service: tts <-> twitch integration"
}
