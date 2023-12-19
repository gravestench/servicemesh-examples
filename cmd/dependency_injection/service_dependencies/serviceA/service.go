package serviceA

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/gravestench/servicemesh"
)

type hasB interface{ B() string }

func New(name string) *Service {
	return &Service{
		name:         name,
		creationTime: time.Now(),
		readyDelay:   time.Duration(rand.Intn(5)+2) * time.Second,
	}
}

type Service struct {
	log  *slog.Logger
	name string

	creationTime time.Time
	readyDelay   time.Duration
	dependency   hasB // depends on service B
}

func (s *Service) A() string {
	return "this message came from ServiceA"
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.log.Info("calling B()", "message from B", s.dependency.B())
	return
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Ready() bool {
	return time.Since(s.creationTime) > s.readyDelay
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.log = logger
}

func (s *Service) DependenciesResolved() bool {
	return s.dependency != nil
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if b, ok := service.(hasB); ok {
			s.dependency = b // If we find our hasB, we assign it!
			break
		}
	}
}
