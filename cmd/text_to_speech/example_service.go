package main

import (
	"math/rand"
	"time"

	"github.com/gravestench/servicesmesh-examples/services/text_to_speech"
)

type exampleServiceThatUsesTts struct {
	tts text_to_speech.ConvertsTextToSpeech
}

func (s *exampleServiceThatUsesTts) DependenciesResolved() bool {
	return s.tts != nil
}

func (s *exampleServiceThatUsesTts) ResolveDependencies(mesh servicemesh.M) {
	for _, service := range servicemesh.Services() {
		if candidate, ok := service.(text_to_speech.ConvertsTextToSpeech); ok {
			s.tts = candidate
		}
	}
}

func (s *exampleServiceThatUsesTts) Init(mesh servicemesh.M) {
	go s.loopSpeakRandomStuff()
}

func (s *exampleServiceThatUsesTts) loopSpeakRandomStuff() {
	voices := s.tts.Voices()
	for {
		randomVoice := voices[rand.Intn(len(voices))]
		s.tts.SetVoice(randomVoice)
		s.tts.Speak(generateRandomPhrase())
	}
}

func (s *exampleServiceThatUsesTts) Name() string {
	return "example service that uses text to speech"
}

var (
	adjectives = []string{
		"happy", "sad", "brave", "kind", "smart", "funny", "silly",
		"crazy", "friendly", "honest", "curious", "energetic",
		"thoughtful", "creative", "patient", "generous",
	}

	nouns = []string{
		"dog", "cat", "book", "friend", "world", "car", "sun",
		"moon", "flower", "tree", "house", "coffee", "song",
		"smile", "dream", "mountain", "river", "ocean", "cloud",
	}

	verbs = []string{
		"run", "jump", "play", "sing", "dance", "read", "write",
		"learn", "explore", "create", "help", "love", "laugh",
		"think", "smile", "dream", "inspire", "imagine", "enjoy",
	}
)

func generateRandomPhrase() string {
	rand.Seed(time.Now().UnixNano())

	phrase := getRandomElement(adjectives) + " " +
		getRandomElement(nouns) + " " +
		getRandomElement(verbs) + " " +
		getRandomElement(adjectives) + " " +
		getRandomElement(nouns) + "."

	return phrase
}

func getRandomElement(list []string) string {
	index := rand.Intn(len(list))
	return list[index]
}
