package serviceB

import (
	"log/slog"

	"github.com/gravestench/servicemesh"
)

type hasA interface{ A() string }

func New(name string) *Service {
	return &Service{
		name: name,
	}
}

type Service struct {
	log  *slog.Logger
	name string

	dependency hasA // depends on service B
}

func (s *Service) B() string {
	return "this message came from ServiceB"
}

func (s *Service) Init(mesh servicemesh.M) {
	s.log.Info("calling A()", "message from A", s.dependency.A())
	return
}

func (s *Service) Name() string {
	return s.name
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

func (s *Service) ResolveDependencies(mesh servicemesh.M) {
	// here, we iterate over all services from the runtime
	// and check if the service implements something we need.
	for _, service := range mesh.Services() {
		if a, ok := service.(hasA); ok {
			s.dependency = a // If we find our hasA, we assign it!
		}
	}
}
