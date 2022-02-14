package support

import (
	"crypto/tls"
	"errors"
	"net"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-mail/mail"
	"github.com/bitwormhole/starter/markup"
)

// SMTPSender 默认的邮件发送组件
type SMTPSender struct {
	markup.Component `id:"mail.Sender"`

	NeedWorkAround bool `inject:"${mail.smtp.workaround}"`

	SenderAddress string `inject:"${mail.sender.address}"`
	SenderName    string `inject:"${mail.sender.name}"`

	SMTPHost     string `inject:"${mail.smtp.host}"`
	SMTPPort     int    `inject:"${mail.smtp.port}"`
	SMTPUser     string `inject:"${mail.smtp.user}"`
	SMTPPassword string `inject:"${mail.smtp.password}"`
}

func (inst *SMTPSender) _Impl() mail.Sender {
	return inst
}

// Send 发送邮件
func (inst *SMTPSender) Send(m *mail.Mail) error {

	const defaultPort = 25
	host := inst.SMTPHost
	port := inst.SMTPPort
	from := m.FromAddr
	to := m.ToAddr

	if port < 1 {
		port = defaultPort
	}

	hostPort := host + ":" + strconv.Itoa(port)
	auth := inst.prepareAuth(m)
	msg := inst.prepareMessage(m)

	if inst.NeedWorkAround {
		workaround := smtpSendMailWorkaround{host: host}
		return workaround.SendMail(hostPort, auth, from, to, msg)
	}

	return smtp.SendMail(hostPort, auth, from, to, msg)
}

func (inst *SMTPSender) prepareAuth(m *mail.Mail) smtp.Auth {
	ident := ""
	username := inst.SMTPUser
	password := inst.SMTPPassword
	host := inst.SMTPHost
	return smtp.PlainAuth(ident, username, password, host)
}

func (inst *SMTPSender) prepareMessage(m *mail.Mail) []byte {

	const nl = "\r\n"
	contentType := inst.prepareContentType(m)
	tolist := inst.stringifyToAddrList(m)
	builder := strings.Builder{}

	builder.WriteString("To: ")
	builder.WriteString(tolist)
	builder.WriteString(nl)

	builder.WriteString("From: ")
	builder.WriteString(m.FromAddr)
	builder.WriteString(nl)

	builder.WriteString("Subject: ")
	builder.WriteString(m.Title)
	builder.WriteString(nl)

	builder.WriteString("Content-Type: ")
	builder.WriteString(contentType)
	builder.WriteString(nl)

	builder.WriteString(nl)
	builder.WriteString(m.Content)

	str := builder.String()
	return []byte(str)
}

func (inst *SMTPSender) prepareContentType(m *mail.Mail) string {
	t := strings.ToLower(m.ContentType)
	if strings.Contains(t, "html") {
		t = "text/html"
	} else {
		t = "text/plain"
	}
	return t + "; charset=UTF-8"
}

func (inst *SMTPSender) stringifyToAddrList(m *mail.Mail) string {
	sep := ""
	builder := strings.Builder{}
	list := m.ToAddr
	for _, item := range list {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(item)
		sep = ";"
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

// HookStartTLS nil, except for tests
type HookStartTLS func(*tls.Config)

type smtpSendMailWorkaround struct {
	testHookStartTLS HookStartTLS
	host             string
}

func (inst *smtpSendMailWorkaround) _Impl() {}

func (inst *smtpSendMailWorkaround) validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("smtp: A line must not contain CR or LF")
	}
	return nil
}

func (inst *smtpSendMailWorkaround) Dial(addr string, cfg *tls.Config) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func (inst *smtpSendMailWorkaround) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {

	if err := inst.validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err := inst.validateLine(recp); err != nil {
			return err
		}
	}
	config := &tls.Config{ServerName: inst.serverName()}
	c, err := inst.Dial(addr, config)
	if err != nil {
		return err
	}
	defer c.Close()

	if a != nil {
		if err = c.Auth(a); err != nil {
			return err
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// -> c.serverName
func (inst *smtpSendMailWorkaround) serverName() string {
	return inst.host
}
