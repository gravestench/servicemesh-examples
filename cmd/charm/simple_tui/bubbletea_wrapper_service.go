package main

import (
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gravestench/servicemesh"
)

type serviceWithModel interface {
	Model() tea.Model
}

type bubbleteaService struct {
	*tea.Program
	mux sync.Mutex
}

func (s *bubbleteaService) Ready() bool { return true }

func (b *bubbleteaService) Init(mesh servicemesh.Mesh) {
	go b.runLoop()

	b.bindExisting(mesh.Services())

	mesh.Events().On(servicemesh.EventServiceAdded, func(...any) {
		b.bindExisting(mesh.Services())
	})
}

func (b *bubbleteaService) runLoop() {
	for {
		time.Sleep(time.Second)

		if b.Program == nil {
			continue
		}

		b.Program.Run()
	}
}

func (b *bubbleteaService) bindExisting(services []servicemesh.Service) {
	var models []tea.Model

	for _, service := range services {
		if candidate, ok := service.(serviceWithModel); ok {
			models = append(models, candidate.Model())
		}
	}

	if len(models) < 1 {
		return
	}

	if b.Program != nil {
		b.Program.Kill()
	}

	b.Program = tea.NewProgram(b.newMainModel(models))
}

func (b *bubbleteaService) newMainModel(models []tea.Model) tea.Model {
	return &mainModel{models: models}
}

func (b *bubbleteaService) Name() string {
	return "Bubbletea"
}

type mainModel struct {
	models            []tea.Model
	currentModelIndex int
}

func (m mainModel) Init() (cmd tea.Cmd) {
	for _, model := range m.models {
		cmd = model.Init()
	}

	return cmd
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.models) < 1 {
		return m, nil
	}

	if m.currentModelIndex < 0 || m.currentModelIndex >= len(m.models) {
		m.currentModelIndex = 0
	}

	return m.models[m.currentModelIndex].Update(msg)
}

func (m mainModel) View() string {
	if len(m.models) < 1 {
		return ""
	}

	if m.currentModelIndex < 0 || m.currentModelIndex >= len(m.models) {
		m.currentModelIndex = 0
	}

	return m.models[m.currentModelIndex].View()
}
