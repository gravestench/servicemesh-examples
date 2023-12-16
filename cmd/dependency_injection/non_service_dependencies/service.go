package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gravestench/runtime/pkg"
	"github.com/gravestench/servicemesh"
)

func newServiceWithAsyncDependencyResolution() *Service {
	s := &Service{}
	s.uuid = uuid.New()

	// this is the example of our dependency being resolved. This could be
	// something like waiting for a SQL database connection to get set up,
	// or a file from the web being pulled, or user input, whatever.
	go func() {
		time.Sleep(generateRandomDuration(10) + (time.Second * 3))
		s.dependency = "resolved" // now it isn't nil
	}()

	return s
}

func generateRandomDuration(maxSeconds int) time.Duration {
	return time.Duration(rand.Intn(maxSeconds)) * time.Second
}

type Service struct {
	// this is the simplest example of a dependency. Below,
	// in DependenciesResolved we check if it is nil
	dependency any

	uuid uuid.UUID    // just for identification in the logs, not really part of this example
	log  *slog.Logger // also not really part of this DI example
}

func (s *Service) DependenciesResolved() bool {
	return s.dependency != nil
}

func (s *Service) ResolveDependencies(_ pkg.IsRuntime) {
	// in this example, we are not using the runtime to find our dependencies,
	// they are resolved some other way (this is up to you). However, we
	// do implement servicemesh.HasDependencies so that the runtime knows not
	// to call Init unless until we have resolved our deps ourselves.
}

func (s *Service) Init(mesh servicemesh.M) {
	return
}

func (s *Service) Name() string {
	return fmt.Sprintf("Service (ID: %.5s)", s.uuid.String())
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.log = logger
}
