package auth

import (
	"context"
	"time"

	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateLogout(input string) (nextState, message string, err error) {
	switch input {
	case constants.CmdEnter:
		timeout := 10 * time.Second
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		defer cancelFunc()

		h.authService.Logout(ctx)

		h.typing = ""
		h.userAuthed = false
		h.position = 0
		return constants.StateMainMenu, "Вы успешно вышли из аккаунта.", nil
	case constants.CmdBack:
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	default:
		return constants.StateLogout, "Нажмите Enter для выхода или Back для возврата в главное меню.", nil
	}
}
