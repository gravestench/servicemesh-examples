package main

type receiver struct {
	logger *slog.Logger
}

func (r *receiver) SetLogger(logger *slog.Logger) {
	r.logger = logger
}

func (r *receiver) Logger() *slog.Logger {
	return r.logger
}

func (r *receiver) Init(mesh servicemesh.M) {
	rt.Events().On("test", func(args ...any) {
		r.logger.Info().Msgf("event 'test' recieved, args are: %+v", args)
	})
}

func (r *receiver) Name() string {
	return "receiver"
}
