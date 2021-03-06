package mail

import (
	"errors"
	"net/smtp"
)

type Authentication struct {
	username, password string
}

func NewAuthentication(username, password string) smtp.Auth {
	return &Authentication{username, password}
}

func (a *Authentication) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *Authentication) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			{
				return []byte(a.username), nil
			}
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}
