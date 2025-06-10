package render

import (
	"os"

	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateQuit(input string) (nextState, message string, err error) {
	switch input {
	case constants.CmdEnter:
		os.Exit(0)
	case constants.CmdBack:
		return constants.StateMainMenu, h.RenderMenu(), nil
	default:
		return constants.StateQuit, "Нажмите Enter для выхода или Back для возврата в главное меню.", nil
	}

	return constants.StateQuit, "Выход из приложения...", nil
}
