package auth

import (
	"os"

	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateQuit(input string) (nextState, message string, err error) {
	switch input {
	case constants.CmdEnter:
		h.typing = ""
		os.Exit(0)
	case constants.CmdBack:
		h.typing = ""
		return constants.StateMainMenu, h.RenderMenu(), nil
	default:
		h.typing = ""
		return constants.StateQuit, "Нажмите Enter для выхода или Back для возврата в главное меню.", nil
	}

	return constants.StateQuit, "Выход из приложения...", nil
}
