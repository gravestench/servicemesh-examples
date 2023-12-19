package serviceB

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/gravestench/servicemesh"
)

type hasA interface{ A() string }

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

	dependency hasA // depends on service B
}

func (s *Service) B() string {
	return "this message came from ServiceB"
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.log.Info("calling A()", "message from A", s.dependency.A())
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
	// here, we iterate over all services from the service mesh
	// and check if the service implements something we need.
	for _, service := range services {
		if a, ok := service.(hasA); ok {
			s.dependency = a // If we find our hasA, we assign it!
		}
	}
}
