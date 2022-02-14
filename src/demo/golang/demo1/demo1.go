package demo1

import (
	"github.com/bitwormhole/starter-mail/mail"
	"github.com/bitwormhole/starter/markup"
)

type Demo1 struct {
	markup.Component `initMethod:"Init"`

	//  AppContext application.Context `inject:"context"`

	FromAddr string `inject:"${demo.send.from.addr}"`
	ToAddr   string `inject:"${demo.send.to.addr}"`

	Sender mail.Sender `inject:"#mail.Sender"`
}

func (inst *Demo1) Init() error {
	return inst.doSend()
}

func (inst *Demo1) doSend() error {

	m := &mail.Mail{}
	to := inst.ToAddr
	from := inst.FromAddr

	m.FromAddr = from
	m.ToAddr = []string{to}
	m.Title = "demo"
	m.ContentType = "html"
	m.Content = "<html><body><h1>demo</h1></body></html>"

	return inst.Sender.Send(m)
}
