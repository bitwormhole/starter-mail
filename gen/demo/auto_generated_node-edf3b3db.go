// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package demo

import (
	mail0xcd88fb "github.com/bitwormhole/starter-mail/mail"
	demo10x334aa8 "github.com/bitwormhole/starter-mail/src/demo/golang/demo1"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDemo1 struct {
	instance *demo10x334aa8.Demo1
	 markup0x23084a.Component `initMethod:"Init"`
	FromAddr string `inject:"${demo.send.from.addr}"`
	ToAddr string `inject:"${demo.send.to.addr}"`
	Sender mail0xcd88fb.Sender `inject:"#mail.Sender"`
}

