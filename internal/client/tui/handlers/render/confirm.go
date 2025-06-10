package render

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/tui/constants"
)

func (h *handler) stateOtpRequested(input string) (nextState, message string, err error) {
	switch {
	case input == constants.CmdEnter:
		if len(h.otp) != 6 {
			return constants.StateOtpRequested, "OTP должен состоять из 6 цифр. Пожалуйста, введите код:", nil
		}

		timeout := 10 * time.Second // время ожидания для подтверждения OTP
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		defer cancelFunc()

		err := h.authService.ConfirmOTP(ctx, string(h.otp))
		if err != nil {
			switch {
			case errors.Is(err, errorx.ErrOtpExpired):
				h.otp = nil
				h.isLogin = true
				return constants.StateLoginPassword, "OTP истек. Пожалуйста, запросите новый код.", nil
			case errors.Is(err, errorx.ErrOtpInvalid):
				h.otp = nil
				return constants.StateOtpRequested, "Неверный OTP. Пожалуйста, введите правильный код.", nil
			case errors.Is(err, errorx.ErrUserNotFound):
				h.otp = nil
				h.isLogin = false
				return constants.StateLoginPassword, "Пользователь не найден. Пожалуйста, зарегистрируйтесь.", nil
			default:
				return constants.StateOtpRequested, "Ошибка подтверждения OTP: " + err.Error(),
					fmt.Errorf("otp validation error: %w", err)
			}
		}

		h.otp = nil
		h.userAuthed = true
		h.position = 0
		if !h.isLogin {
			return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
		}

		return constants.StateAuthorizedMainMenu, h.RenderMenu(), nil
	case input == constants.CmdBack && len(h.otp) > 0:
		h.otp = h.otp[:len(h.otp)-1]
	case input != "" && input != constants.CmdBack && !ignoreInput[input]:
		h.otp = append(h.otp, []rune(input)...)
	}

	return constants.StateOtpRequested, "Введите OTP: " + string(h.otp), nil
}
