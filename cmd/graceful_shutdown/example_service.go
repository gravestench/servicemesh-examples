package main

type example struct {
	l    *slog.Logger
	name string
}

func (e *example) SetLogger(logger *slog.Logger) {
	e.l = logger
}

func (e *example) Logger() *slog.Logger {
	return e.l
}

func (e *example) Init(r servicemesh.Runtime) {
	return
}

func (e *example) Name() string {
	return e.name
}

func (e *example) OnShutdown() {
	e.l.Info().Msg("doing cleanup here...")
}
