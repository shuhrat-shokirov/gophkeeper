package emailtotp

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"gopkg.in/gomail.v2"

	"gophkeeper/pkg/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
}

type Gateway interface {
	SendEmail(_ context.Context, request *Request) error
}

type gateway struct {
	mail     string
	password string
}

func New(p Params) Gateway {
	return &gateway{
		mail:     p.Config.GetString("email.mail"),
		password: p.Config.GetString("email.password"),
	}
}

func (g *gateway) SendEmail(_ context.Context, request *Request) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.mail)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/plain", request.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, g.mail, g.password)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send email fail: %v", err)
	}

	return nil
}
