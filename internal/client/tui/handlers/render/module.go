package render

import (
	"context"
	"os"

	"go.uber.org/fx"

	"gophkeeper/internal/client/services/auth"
	"gophkeeper/internal/client/services/data"
	"gophkeeper/internal/client/tui/constants"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	AuthService auth.Service
	DataService data.Service
}

type Handler interface {
	HandleInput(state, input string) (newState, screen string, err error)
	RenderMenu() string
	CheckAuth() bool
}

type handler struct {
	email    []rune
	password []rune
	otp      []rune

	addTitle    []rune
	addLogin    []rune
	addPassword []rune
	addNote     []rune
	addTextData []rune
	addCardPan  []rune
	addCardCvv  []rune
	addCardExp  []rune
	addFilePath []rune

	isAddLogin bool
	isAddText  bool
	isAddFile  bool
	isAddCard  bool

	position   int // для навигации вверх/вниз
	isLogin    bool
	userAuthed bool // флаг, что пользователь авторизован

	authService auth.Service
	dataService data.Service
}

func New(p Params) Handler {
	return &handler{
		authService: p.AuthService,
		dataService: p.DataService,
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

	case constants.StateAddData:
		return h.stateAddData(input)
	case constants.StateAddTitleData:
		return h.stateAddTitleData(input)
	case constants.StateAddLoginData:
		return h.stateAddLoginData(input)
	case constants.StateAddPasswordData:
		return h.stateAddPasswordData(input)
	case constants.StateAddNoteData:
		return h.stateAddNoteData(input)
	case constants.StateSuccessAddData:
		return h.stateSuccessAddData(input)

	case constants.StateAddTextData:
		return h.stateAddTextData(input)

	case constants.StateAddCardPan:
		return h.stateAddCardPanData(input)
	case constants.StateAddCardCvv:
		return h.stateAddCardCvvData(input)
	case constants.StateAddCardExpiry:
		return h.stateAddCardExpiryData(input)

	case constants.StateAddFileData:
		return h.stateAddFileData(input)
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
