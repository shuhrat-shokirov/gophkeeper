package auth

import (
	"gophkeeper/internal/client/tui/constants"
	"gophkeeper/pkg/utils"
)

func (h *handler) stateLoginEmail(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if !utils.ValidateEmail(string(h.email)) {
			return constants.StateLoginEmail,
				"Email не соответствует формату. Пожалуйста, введите правильный email:", nil
		}
		h.typing = "password"
		return constants.StateLoginPassword, "Введите пароль:", nil
	case input == constants.CmdBack && len(h.email) > 0:
		h.email = h.email[:len(h.email)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.email = append(h.email, []rune(input)...)
	}

	return constants.StateLoginEmail, "Введите email: " + string(h.email), nil
}
