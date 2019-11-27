package pinbox

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

type SMTP struct {
	server   string
	username string
	password string
	port     int
}

func CreateSMTP() *SMTP {
	return &SMTP{
		server:   "test.com",
		username: "test",
		password: "test",
		port:     25,
	}
}

func (mta *SMTP) Send(msg MessageContent) error {
	auth := smtp.PlainAuth("", mta.username, mta.password, mta.server)

	host := fmt.Sprintf("%s:%d", mta.server, mta.port)
	body, err := base64.StdEncoding.DecodeString(msg.Content)

	if err != nil {
		return err
	}

	err = smtp.SendMail(host, auth, msg.Author, msg.To, body)

	if err != nil {
		return err
	}

	return nil
}
