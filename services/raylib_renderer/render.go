package raylib_renderer

import (
	"os"
	"time"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/servicemesh"
)

func (s *Service) initRenderer() {
	cfg, err := s.Config()
	if err != nil {
		s.log.Error("getting config", "error", err)
	}

	windowSettings := cfg.Group(keyGroupWindowSettings)
	screenWidth := windowSettings.GetInt(keyScreenWidth)
	screenHeight := windowSettings.GetInt(keyScreenHeight)

	rl.SetTraceLogCallback(func(_ int, message string) {
		s.Logger().Info(message)
	})

	// this requires servicemesh.Run() to be passed into mainthread.Run
	mainthread.Call(func() {
		rl.SetTargetFPS(60)
		rl.InitWindow(int32(screenWidth), int32(screenHeight), "runtime - raylib example")
	})
}

func (s *Service) gatherLayers(mesh servicemesh.M) {
	for {
		for _, service := range mesh.Services() {
			if _, alreadyBound := s.layers[service.Name()]; alreadyBound {
				continue
			}

			if candidate, ok := service.(RenderableLayer); ok {
				s.AddLayer(candidate)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) renderServicesAsLayers(mesh servicemesh.M) {
	s.mux.Lock()
	defer s.mux.Unlock()

	defer func() { _ = recover() /* don't worry about it */ }()

	mainthread.Call(func() {
		defer func() { _ = recover() /* don't worry about it */ }()

		for !rl.WindowShouldClose() {
			rl.BeginMode2D(rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1))
			rl.BeginDrawing()
			rl.SetBlendMode(int32(rl.BlendAlpha))
			rl.ClearBackground(rl.Black)

			for _, name := range s.order {
				if layer := s.layers[name]; layer != nil {
					layer.OnRender()
				}
			}

			rl.EndDrawing()
		}

		rl.CloseWindow()
		s.log.Info("closing window")
		os.Exit(0)
	})
}
