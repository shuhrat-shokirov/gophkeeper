package render

import (
	"time"

	"gophkeeper/internal/client/tui/constants"
)

const (
	limit   = 5
	timeout = 10 * time.Second
)

func (h *handler) stateGetDataList(input string) (nextState, message string, err error) {
	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + 4) % 4
	case constants.CmdDown:
		h.position = (h.position + 1) % 4
	case constants.CmdBack:
		h.position = 0
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	case constants.CmdEnter:
		switch h.position {
		case 0:
			h.position = 0
			h.offset = 0
			h.flushMapID()
			return h.stateLoginList(input)
		case 3:
			h.position = 0
			h.offset = 0
			h.flushMapID()
			return h.stateCardList(input)
		}
	default:
		return constants.StateGetDataList, "Неверная команда. Используйте стрелки для навигации.", nil
	}

	return constants.StateGetDataList, h.renderGetDataList(), nil
}

func (h *handler) renderGetDataList() string {
	options := []string{
		"Показать сохраненные логины",
		"Показать сохраненные тексты",
		"Показать сохраненные файлы",
		"Показать сохраненные карты",
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
