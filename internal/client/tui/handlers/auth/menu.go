package auth

import "gophkeeper/internal/client/tui/constants"

func (h *handler) RenderMenu() string {
	options := []string{"Логин", "Регистрация", "Выход"}

	if h.userAuthed {
		options = []string{"Показать данные", "Добавить данные", "Выйти"}
	}

	screen := "Выберите опцию (стрелки вверх/вниз + Enter):\n"
	for i, opt := range options {
		if i == h.position {
			screen += "> " + opt + "\n"
		} else {
			screen += "  " + opt + "\n"
		}
	}
	return screen
}

func (h *handler) stateMainMenu(input string) (state, message string, err error) {
	h.typing = ""
	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + 3) % 3
	case constants.CmdDown:
		h.position = (h.position + 1) % 3
	case constants.CmdEnter:
		switch h.position {
		case 0: // Логин
			h.typing = "email"
			h.isLogin = true
			return constants.StateLoginEmail, "Введите email:", nil
		case 1: // Регистрация
			h.typing = "email"
			h.isLogin = false
			return constants.StateLoginEmail, "Введите email для регистрации:", nil
		case 2: // Выход
			return constants.StateQuit, "Выход...", nil
		}
	}

	return constants.StateMainMenu, h.RenderMenu(), nil
}
