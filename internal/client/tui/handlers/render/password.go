package render

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateLoginPassword(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.password) < 6 {
			return constants.StateLoginPassword,
				"Пароль не может быть меньше 6 символов. Пожалуйста, введите пароль:", nil
		}

		timeout := 10 * time.Second // время ожидания для регистрации
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		defer cancelFunc()

		if !h.isLogin {
			err := h.authService.Register(ctx, string(h.email), string(h.password))
			if err != nil {
				if errors.Is(err, errorx.ErrUserAlreadyExists) {
					// Если пользователь уже существует, то предлагаем ввести пароль для входа
					h.isLogin = true
					h.password = nil // очищаем пароль для нового ввода
					return constants.StateLoginPassword, "Пользователь с таким email уже существует. Введите пароль:", nil
				}

				return constants.StateLoginPassword, "Ошибка регистрации: " + err.Error(),
					fmt.Errorf("failed to register: %w", err)
			}

			h.password = nil
			return constants.StateOtpRequested, "OTP отправлен на почту. Введите код:", nil
		}

		err := h.authService.Login(ctx, string(h.email), string(h.password))
		if err != nil {
			if errors.Is(err, errorx.ErrInvalidCredentials) {
				return constants.StateLoginPassword, "Неверный email или пароль. Пожалуйста, попробуйте снова:", nil
			}

			return constants.StateLoginPassword, "Ошибка входа: " + err.Error(),
				fmt.Errorf("failed to login: %w", err)
		}

		h.isLogin = true
		h.password = nil
		return constants.StateOtpRequested, "OTP отправлен на почту. Введите код:", nil

	case input == constants.CmdBack && len(h.password) > 0:
		h.password = h.password[:len(h.password)-1]

	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.password = append(h.password, []rune(input)...)
	}

	return constants.StateLoginPassword, "Введите пароль: " + string(h.password), nil
}
