package cli_flags

import (
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gravestench/servicemesh"
)

// Service is responsible for iterating over other services and
// filtering out the CLI args which don't apply to that particular service.
// this addresses an issue with the golang flag implementation that
// prevents flags from being defined per-module. If you try to parse
// flags which are not supposed to be parsed, it errors early and doesn't
// parse all the flags.
type Service struct {
	log  *slog.Logger
	args []string

	parsed map[string]struct{}
	mux    sync.Mutex
}

func (s *Service) SetLogger(l *slog.Logger) {
	s.log = l
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func (s *Service) Init(mesh servicemesh.M) {
	s.args = make([]string, len(os.Args))
	copy(s.args, os.Args)

	s.mux.Lock()
	s.parsed = make(map[string]struct{})
	s.mux.Unlock()

	go s.loopApplyFlags(mesh)
}

func (s *Service) Name() string {
	return "CLI Flags"
}

func (s *Service) loopApplyFlags(manager servicemesh.M) {
	for {
		for _, candidate := range manager.Services() {
			s.applyFlags(candidate)
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) applyFlags(candidate servicemesh.Service) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if svc, ok := candidate.(ServiceThatUsesFlags); ok {
		if _, found := s.parsed[svc.Name()]; found {
			return
		}

		args := s.getArgs()
		args = s.filterArgsForService(svc, args)

		s.log.Info("parsing CLI flags", "service", svc.Name())

		flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		s.parsed[svc.Name()] = struct{}{}

		if err := svc.Parse(flagSet, args); err != nil {
			return
		}
	}
}

func (s *Service) getArgs() []string {
	args := make([]string, len(s.args))
	copy(args, s.args)

	return args
}

func (s *Service) filterArgsForService(fs ServiceThatUsesFlags, args []string) []string {
	matchArg := `--?[a-zA-Z0-9-]+( [^-]?[^ ]*)?`
	m := regexp.MustCompile(matchArg)
	args = m.FindAllString(strings.Join(args, " "), -1)
	filtered := make([]string, 0)

	for _, arg := range args {
		flagsTheServicesLooksFor := fs.Flags()
		flagsTheServicesLooksFor = append(flagsTheServicesLooksFor, "--help")
		if containsOneOfTheFlags(arg, flagsTheServicesLooksFor) {
			filtered = append(filtered, strings.Split(arg, " ")...)
		}
	}

	return filtered
}

func containsOneOfTheFlags(arg string, argsNeededByService []string) bool {
	for _, n := range argsNeededByService {
		if strings.Contains(arg, n) {
			return true
		}
	}

	return false
}
