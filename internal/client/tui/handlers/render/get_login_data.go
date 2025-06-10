//nolint:dupl,gocritic
package render

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) renderLoginList() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	list, err := h.dataService.GetLoginList(ctx, limit, h.offset)
	if err != nil {
		return "", fmt.Errorf("ошибка получения списка логинов: %w", err)
	}

	var options = make([]string, 0, len(list))
	for i, item := range list {
		if item.Title == "" {
			item.Title = "Без названия"
		}
		options = append(options, fmt.Sprintf(" %d. %s", int(h.offset)+len(options)+1, item.Title))
		h.listMapID[i] = item.ID
	}

	screen := fmt.Sprintf(`Список логинов: 
Страница: %d
Выберите опцию (стрелки вверх/вниз + Enter):
`, int(h.offset)/limit+1)

	for i, opt := range options {
		if i == h.position {
			screen += "> " + opt + "\n"
		} else {
			screen += "  " + opt + "\n"
		}
	}

	screen += "\n "

	if h.offset > 0 {
		screen += "[p] - Предыдущая страница "
	}

	if len(options) == limit {
		screen += "[n] - Следующая страница "
	}

	return screen, nil
}

func (h *handler) stateLoginList(input string) (nextState, message string, err error) {
	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + h.listCount) % h.listCount
	case constants.CmdDown:
		h.position = (h.position + 1) % h.listCount
	case constants.CmdBack:
		h.position = 0
		return constants.StateGetDataList, h.renderGetDataList(), nil
	case constants.CmdNextPage:
		h.offset += limit
		h.flushMapID()
	case constants.CmdPrevPage:
		h.offset -= limit
		h.flushMapID()
	case constants.CmdEnter:
		if len(h.listMapID) != 0 {
			id, ok := h.listMapID[h.position]
			if !ok {
				return constants.StateLoginList, "Ошибка: ID логина не найден.", nil
			}

			screen, err := h.renderLoginByID(id)
			if err != nil {
				if errors.Is(err, errorx.ErrNotFound) {
					return constants.StateLoginList, "Логин не найден.", nil
				}

				return constants.StateLoginList, "Ошибка получения логина: " + err.Error(), err
			}

			return constants.StateLoginList, screen, nil
		}
	default:
		return constants.StateLoginList, "Неверная команда. Используйте стрелки для навигации.", nil
	}

	result, err := h.renderLoginList()
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return constants.StateLoginList, `Больше нет данных для отображения.`, nil
		}
		return constants.StateLoginList, "Ошибка получения списка логинов: " + err.Error(), err
	}

	h.listCount = len(h.listMapID)

	return constants.StateLoginList, result, nil
}

func (h *handler) renderLoginByID(id int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	loginData, err := h.dataService.GetLoginByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("ошибка получения логина по ID: %w", err)
	}

	screen := fmt.Sprintf(`Описание: %s
Логин: %s
Пароль: %s
Заметка: %s
Последнее обновление: %s
`, loginData.Title, loginData.Login,
		loginData.Pass, loginData.Note,
		loginData.UpdatedAt.Format(time.DateTime))

	return screen, nil
}
