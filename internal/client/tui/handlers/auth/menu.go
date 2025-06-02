package auth

import (
	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) RenderMenu() string {
	options := []string{"Логин", "Регистрация", "Завершить работу"}

	if h.userAuthed {
		options = []string{"Показать данные", "Добавить данные", "Выход из аккаунта", "Завершить работу"}
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
	positionCount := 3 // Количество позиций в главном меню
	if h.userAuthed {
		positionCount = 4 // Количество позиций в меню авторизованного пользователя
	}

	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + positionCount) % positionCount
	case constants.CmdDown:
		h.position = (h.position + 1) % positionCount
	case constants.CmdEnter:
		if h.userAuthed {
			return h.stateAuthorizedMainMenu(input)
		}
		switch h.position {
		case 0: // Логин
			h.typing = "email"
			h.isLogin = true
			return constants.StateLoginEmail, "Введите email:", nil
		case 1: // Регистрация
			h.typing = "email"
			h.isLogin = false
			return constants.StateLoginEmail, "Введите email для регистрации:", nil
		case 2: // Завершить работу
			return constants.StateQuit, "Завершение работы приложения...", nil
		default:
			h.typing = ""
			h.position = 0 // Сброс позиции при возврате в главное меню
			return constants.StateMainMenu, h.RenderMenu(), nil
		}
	case constants.CmdBack:
		h.typing = ""
		h.position = 0 // Сброс позиции при возврате в главное меню
		return constants.StateMainMenu, h.RenderMenu(), nil
	}

	return constants.StateMainMenu, h.RenderMenu(), nil
}

func (h *handler) stateAuthorizedMainMenu(input string) (state, message string, err error) {
	h.typing = ""
	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + 4) % 4
	case constants.CmdDown:
		h.position = (h.position + 1) % 4
	case constants.CmdEnter:
		switch h.position {
		case 0: // Показать данные
			return constants.StateDataMenu, "Показать данные", nil
		case 1: // Добавить данные
			return constants.StateSwitchToData, "Добавить данные", nil
		case 2: // Выйти
			return constants.StateLogout, "Нажмите Enter для выхода или Back для возврата в главное меню.", nil
		case 3: // Завершить работу
			return constants.StateQuit, "Завершение работы приложения...", nil
		default:
			h.typing = ""
			h.position = 0 // Сброс позиции при возврате в главное меню
			return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
		}
	default:
		h.typing = ""
		h.position = 0 // Сброс позиции при возврате в главное меню
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	}

	return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil

}
