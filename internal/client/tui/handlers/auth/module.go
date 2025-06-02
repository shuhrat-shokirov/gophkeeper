package auth

import (
	"context"
	"os"

	"go.uber.org/fx"

	"gophkeeper/internal/client/services/auth"
	"gophkeeper/internal/client/tui/constants"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	AuthService auth.Service
}

type Handler interface {
	HandleInput(state, input string) (newState, screen string, err error)
	RenderMenu() string
	CheckAuth() bool
}

type handler struct {
	email      []rune
	password   []rune
	otp        []rune
	position   int    // для навигации вверх/вниз
	typing     string // "email", "password", "otp"
	isLogin    bool
	userAuthed bool // флаг, что пользователь авторизован

	authService auth.Service
}

func New(p Params) Handler {
	return &handler{
		authService: p.AuthService,
	}
}

var ignoreInput = map[string]bool{
	constants.CmdUp:   true,
	constants.CmdDown: true,
}

func (h *handler) HandleInput(state, input string) (nextState, screen string, err error) {
	if input == constants.CmdForceQuit {
		os.Exit(0)
		return
	}

	switch state {
	case constants.StateMainMenu:
		return h.stateMainMenu(input)
	case constants.StateAuthorizedMainMenu:
		return h.stateAuthorizedMainMenu(input)
	case constants.StateLoginEmail:
		return h.stateLoginEmail(input)
	case constants.StateLoginPassword:
		return h.stateLoginPassword(input)
	case constants.StateOtpRequested:
		return h.stateOtpRequested(input)
	case constants.StateLogout:
		return h.stateLogout(input)
	case constants.StateQuit:
		return h.stateQuit(input)
	}

	return constants.StateMainMenu, h.RenderMenu(), nil
}

func (h *handler) CheckAuth() bool {
	err := h.authService.CheckAuth(context.Background())
	if err != nil {
		return false
	}

	h.userAuthed = true

	return true
}
