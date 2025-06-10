//nolint:dupl,gocritic
package render

import (
	"context"
	"fmt"
	"os"
	"time"
	"unicode"

	"gophkeeper/internal/client/services/data"
	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateAddData(input string) (nextState, message string, err error) {

	switch input {
	case constants.CmdUp:
		h.position = (h.position - 1 + 4) % 4 // Предполагаем, что в меню добавления данных 3 позиции
	case constants.CmdDown:
		h.position = (h.position + 1) % 4 // Циклический переход по позициям меню добавления данных
	case constants.CmdBack:
		h.position = 0 // Сброс позиции при возврате в главное меню
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	case constants.CmdEnter:
		switch h.position {
		case 0: // Добавить логин/пароль
			h.isAddLogin = true
			return constants.StateAddTitleData, "Введите заголовок: ", nil
		case 1: // Добавить текстовую заметку
			h.isAddText = true
			return constants.StateAddTitleData, "Введите заголовок: ", nil
		case 2: // Добавить банковскую карту
			h.isAddCard = true
			return constants.StateAddTitleData, "Введите заголовок: ", nil
		case 3: // Добавить файл
			h.isAddFile = true
			return constants.StateAddTitleData, "Введите заголовок: ", nil
		default:
			h.position = 0 // Сброс позиции при возврате в главное меню
			h.isLogin = false
			return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
		}
	}

	h.isLogin = false
	return constants.StateAddData, h.renderAddDataMenu(), nil
}

func (h *handler) renderAddDataMenu() string {
	options := []string{
		"Добавить логин/пароль",
		"Добавить текстовую заметку",
		"Добавить банковскую карту",
		"Добавить файл",
	}

	screen := "Выберите тип данных для добавления (стрелки вверх/вниз + Enter):\n"
	for i, opt := range options {
		if i == h.position {
			screen += "> " + opt + "\n"
		} else {
			screen += "  " + opt + "\n"
		}
	}
	return screen
}

func (h *handler) stateAddTitleData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addTitle) == 0 {
			return constants.StateAddTitleData, "Заголовок не может быть пустым. Пожалуйста, введите заголовок:", nil
		}

		switch {
		case h.isAddLogin:
			return constants.StateAddLoginData, "Введите логин: ", nil
		case h.isAddText:
			return constants.StateAddTextData, "Введите текстовые данные: ", nil
		case h.isAddCard:
			return constants.StateAddCardPan, "Введите Номер карты (PAN): ", nil
		case h.isAddFile:
			return constants.StateAddFileData, "Введите путь к файлу: ", nil
		}
	case input == constants.CmdBack && len(h.addTitle) > 0:
		h.addTitle = h.addTitle[:len(h.addTitle)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addTitle = append(h.addTitle, []rune(input)...)
	}

	return constants.StateAddTitleData, "Введите заголовок: " + string(h.addTitle), nil
}

func (h *handler) stateAddLoginData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addLogin) == 0 {
			return constants.StateAddLoginData, "Логин не может быть пустым. Пожалуйста, введите логин:", nil
		}

		return constants.StateAddPasswordData, "Введите пароль:", nil
	case input == constants.CmdBack && len(h.addLogin) > 0:
		h.addLogin = h.addLogin[:len(h.addLogin)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addLogin = append(h.addLogin, []rune(input)...)
	}

	return constants.StateAddLoginData, "Введите логин: " + string(h.addLogin), nil
}

func (h *handler) stateAddPasswordData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addPassword) == 0 {
			return constants.StateAddPasswordData, "Пароль не может быть пустым. Пожалуйста, введите пароль:", nil
		}

		return constants.StateAddNoteData, "Введите метаданные (необязательно): ", nil
	case input == constants.CmdBack && len(h.addPassword) > 0:
		h.addPassword = h.addPassword[:len(h.addPassword)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addPassword = append(h.addPassword, []rune(input)...)
	}

	return constants.StateAddPasswordData, "Введите пароль: " + string(h.addPassword), nil
}

func (h *handler) stateAddNoteData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		switch {
		case h.isAddLogin:
			err := h.dataService.SaveLogin(ctx, &data.LoginData{
				Login: string(h.addLogin),
				Pass:  string(h.addPassword),
				Title: string(h.addTitle),
				Note:  string(h.addNote),
			})
			if err != nil {
				return constants.StateAddNoteData,
					"Ошибка при добавлении данных: " + err.Error(),
					fmt.Errorf("login save error: %w", err)
			}
		case h.isAddText:
			err := h.dataService.SaveText(ctx, &data.TextData{
				Title:   string(h.addTitle),
				Content: string(h.addTextData),
				Note:    string(h.addNote),
			})
			if err != nil {
				return constants.StateAddNoteData,
					"Ошибка при добавлении данных: " + err.Error(),
					fmt.Errorf("text save error: %w", err)
			}
		case h.isAddCard:
			err := h.dataService.SaveCard(ctx, &data.CardData{
				Title:  string(h.addTitle),
				Pan:    string(h.addCardPan),
				Cvv:    string(h.addCardCvv),
				Expiry: string(h.addCardExp),
				Note:   string(h.addNote),
			})
			if err != nil {
				return constants.StateAddNoteData,
					"Ошибка при добавлении данных: " + err.Error(),
					fmt.Errorf("card save error: %w", err)
			}
		case h.isAddFile:
			err := h.dataService.SaveFile(ctx, &data.FileData{
				Title: string(h.addTitle),
				Path:  string(h.addFilePath),
				Note:  string(h.addNote),
			})
			if err != nil {
				return constants.StateAddNoteData,
					"Ошибка при добавлении данных: " + err.Error(),
					fmt.Errorf("file save error: %w", err)
			}
		}

		return constants.StateSuccessAddData, "Данные успешно добавлены! Нажмите Enter для продолжения.", nil
	case input == constants.CmdBack && len(h.addNote) > 0:
		h.addNote = h.addNote[:len(h.addNote)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addNote = append(h.addNote, []rune(input)...)
	}

	return constants.StateAddNoteData, "Введите метаданные: " + string(h.addNote), nil
}

func (h *handler) stateSuccessAddData(input string) (nextState, message string, err error) {
	h.position = 0 // Сброс позиции при возврате в главное меню
	h.isAddText = false
	h.isAddLogin = false
	h.addTitle = nil
	h.addLogin = nil
	h.addPassword = nil
	h.addNote = nil

	switch input {
	case constants.CmdEnter:
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	case constants.CmdBack:
		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	default:
		return constants.StateSuccessAddData, "Данные успешно добавлены! Нажмите Enter для продолжения.", nil
	}
}

func (h *handler) stateAddTextData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addTextData) == 0 {
			return constants.StateAddTextData, "Текст не может быть пустым. Пожалуйста, введите текст:", nil
		}

		return constants.StateAddNoteData, "Введите метаданные (необязательно): ", nil
	case input == constants.CmdBack && len(h.addTextData) > 0:
		h.addTextData = h.addTextData[:len(h.addTextData)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addTextData = append(h.addTextData, []rune(input)...)
	}

	return constants.StateAddTextData, "Введите текстовые данные: " + string(h.addTextData), nil
}

func (h *handler) stateAddCardPanData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addCardPan) == 0 {
			return constants.StateAddCardPan, "PAN не может быть пустым. Пожалуйста, введите PAN:", nil
		}

		if !unicode.IsDigit(h.addCardPan[len(h.addCardPan)-1]) {
			h.addCardPan = nil
			return constants.StateAddCardPan, "PAN должен состоять только из цифр. Пожалуйста, введите корректный PAN:", nil
		}

		if len(h.addCardPan) < 16 || len(h.addCardPan) > 19 {
			return constants.StateAddCardPan, "PAN должен содержать от 16 до 19 цифр. Пожалуйста, введите корректный PAN:", nil
		}

		return constants.StateAddCardCvv, "Введите CVV: ", nil
	case input == constants.CmdBack && len(h.addCardPan) > 0:
		h.addCardPan = h.addCardPan[:len(h.addCardPan)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addCardPan = append(h.addCardPan, []rune(input)...)
	}

	return constants.StateAddCardPan, "Введите Номер карты (PAN): " + string(h.addCardPan), nil
}

func (h *handler) stateAddCardCvvData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addCardCvv) == 0 {
			return constants.StateAddCardCvv, "CVV не может быть пустым. Пожалуйста, введите CVV:", nil
		}

		if !unicode.IsDigit(h.addCardCvv[len(h.addCardCvv)-1]) {
			h.addCardCvv = nil
			return constants.StateAddCardCvv, "CVV должен состоять только из цифр. Пожалуйста, введите корректный CVV:", nil
		}

		if len(h.addCardCvv) < 3 || len(h.addCardCvv) > 4 {
			return constants.StateAddCardCvv, "CVV должен содержать 3 или 4 цифры. Пожалуйста, введите корректный CVV:", nil
		}

		return constants.StateAddCardExpiry, "Введите срок действия карты (MM/YY): ", nil
	case input == constants.CmdBack && len(h.addCardCvv) > 0:
		h.addCardCvv = h.addCardCvv[:len(h.addCardCvv)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addCardCvv = append(h.addCardCvv, []rune(input)...)
	}

	return constants.StateAddCardCvv, "Введите CVV: " + string(h.addCardCvv), nil
}

func (h *handler) stateAddCardExpiryData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addCardExp) == 0 {
			return constants.StateAddCardExpiry,
				"Срок действия карты не может быть пустым. Пожалуйста, введите срок действия:",
				nil
		}

		if !unicode.IsDigit(h.addCardExp[len(h.addCardExp)-1]) && h.addCardExp[len(h.addCardExp)-1] != '/' {
			h.addCardExp = nil
			return constants.StateAddCardExpiry,
				"Срок действия карты должен состоять только из цифр и символа '/'. " +
					"Пожалуйста, введите корректный срок:",
				nil
		}

		if len(h.addCardExp) != 5 || h.addCardExp[2] != '/' {
			return constants.StateAddCardExpiry,
				"Срок действия карты должен быть в формате MM/YY. Пожалуйста, введите корректный срок:",
				nil
		}

		return constants.StateAddNoteData, "Введите метаданные (необязательно): ", nil
	case input == constants.CmdBack && len(h.addCardExp) > 0:
		h.addCardExp = h.addCardExp[:len(h.addCardExp)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addCardExp = append(h.addCardExp, []rune(input)...)
	}

	return constants.StateAddCardExpiry, "Введите срок действия карты (MM/YY): " + string(h.addCardExp), nil
}

// StateAddFileData
func (h *handler) stateAddFileData(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.addFilePath) == 0 {
			return constants.StateAddFileData, "Путь к файлу не может быть пустым. Пожалуйста, введите путь:", nil
		}

		_, err := os.ReadFile(string(h.addFilePath))
		if err != nil {
			return constants.StateAddFileData,
				"Ошибка чтения файла: " + err.Error(),
				fmt.Errorf("file read error: %w", err)
		}

		return constants.StateAddNoteData, "Введите метаданные (необязательно): ", nil
	case input == constants.CmdBack && len(h.addFilePath) > 0:
		h.addFilePath = h.addFilePath[:len(h.addFilePath)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.addFilePath = append(h.addFilePath, []rune(input)...)
	}

	return constants.StateAddFileData, "Введите путь к файлу: " + string(h.addFilePath), nil
}
