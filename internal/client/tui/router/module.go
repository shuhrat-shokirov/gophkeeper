package router

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/fx"

	"gophkeeper/internal/client/tui/constants"
	"gophkeeper/internal/client/tui/handlers/render"
)

var Module = fx.Invoke(New)

type Params struct {
	fx.In

	RenderHandler render.Handler
}

type router struct {
	state    string
	input    string
	screen   string
	quit     bool
	userAuth bool
	handler  func(state, input string) (string, string, error)
}

func New(p Params) {
	module := &router{
		state:   constants.StateMainMenu,
		handler: p.RenderHandler.HandleInput,
		screen:  p.RenderHandler.RenderMenu(),
	}

	if ok := p.RenderHandler.CheckAuth(); ok {
		module.userAuth = true
		module.state = constants.StateAuthorizedMainMenu
		module.screen = p.RenderHandler.RenderMenu()
	}

	program := tea.NewProgram(module)

	if _, err := program.Run(); err != nil {
		log.Printf("error running program: %s", err)
		os.Exit(1)
	}
}

func (r *router) Init() tea.Cmd {
	return nil
}

func (r *router) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if r.quit {
		return r, tea.Quit
	}

	m, ok := msg.(tea.KeyMsg)

	if !ok {
		return r, nil
	}

	key := m.String()

	if key == constants.CmdForceQuit {
		r.quit = true
		return r, tea.Quit
	}

	if key == constants.CmdEnter || key == constants.CmdUp || key == constants.CmdDown || key == constants.CmdBack {
		r.input = key
	} else if len(m.Runes) > 0 {
		r.input = string(m.Runes)
	}

	newState, newScreen, err := r.handler(r.state, r.input)
	if err != nil {
		r.screen = "Ошибка: " + err.Error()
	} else {
		r.state = newState
		r.screen = newScreen
	}

	r.input = ""

	return r, nil
}

func (r *router) View() string {
	return fmt.Sprintf("%s %s", r.screen, r.input)
}
