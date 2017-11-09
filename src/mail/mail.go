package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
)

//func Send(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
//	c, err := smtp.Dial(addr)
//	host, _, _ := net.SplitHostPort(addr)
//	if err != nil {
//		fmt.Println("call dial")
//		return err
//	}
//	defer c.Close()
//
//	if ok, _ := c.Extension("STARTTLS"); ok {
//		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
//		if err = c.StartTLS(config); err != nil {
//			fmt.Println("call start tls")
//			return err
//		}
//	}
//
//	if a != nil {
//		if ok, _ := c.Extension("AUTH"); ok {
//			if err = c.Auth(a); err != nil {
//				fmt.Println("check auth with err:", err)
//				return err
//			}
//		}
//	}
//
//	if err = c.Mail(from); err != nil {
//		return err
//	}
//	for _, addr := range to {
//		if err = c.Rcpt(addr); err != nil {
//			return err
//		}
//	}
//	w, err := c.Data()
//	if err != nil {
//		return err
//	}
//
//	header := make(map[string]string)
//	header["MIME-Version"] = "1.0"
//	header["Content-Type"] = "text/plain; charset=\"utf-8\""
//	header["Content-Transfer-Encoding"] = "base64"
//	message := ""
//	for k, v := range header {
//		message += fmt.Sprintf("%s: %s\r\n", k, v)
//	}
//	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
//	_, err = w.Write([]byte(message))
//
//	if err != nil {
//		return err
//	}
//	err = w.Close()
//	if err != nil {
//		return err
//	}
//	return c.Quit()
//}

type Mail interface {
	Send() error
	From(string)
	To([]string)
	Msg([]byte)
}

type mail struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

func (m *mail) Send() error {
	c, err := smtp.Dial(m.addr)
	host, _, _ := net.SplitHostPort(m.addr)
	if err != nil {
		fmt.Println("call dial")
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			fmt.Println("call start tls")
			return err
		}
	}

	if m.auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(m.auth); err != nil {
				fmt.Println("check auth with err:", err)
				return err
			}
		}
	}

	if err = c.Mail(m.from); err != nil {
		return err
	}
	for _, addr := range m.to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(m.msg)
	_, err = w.Write([]byte(message))

	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func (m *mail) From(from string) {
	m.from = from
}

func (m *mail) To(to []string) {
	m.to = to
}

func (m *mail) Msg(msg []byte) {
	m.msg = msg
}

func New(addr string, auth smtp.Auth) Mail {
	return &mail{addr: addr, auth: auth}
}
