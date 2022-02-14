// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package lib

import (
	support0x5608ca "github.com/bitwormhole/starter-mail/mail/support"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComSMTPSender struct {
	instance *support0x5608ca.SMTPSender
	 markup0x23084a.Component `id:"mail.Sender"`
	NeedWorkAround bool `inject:"${mail.smtp.workaround}"`
	SenderAddress string `inject:"${mail.sender.address}"`
	SenderName string `inject:"${mail.sender.name}"`
	SMTPHost string `inject:"${mail.smtp.host}"`
	SMTPPort int `inject:"${mail.smtp.port}"`
	SMTPUser string `inject:"${mail.smtp.user}"`
	SMTPPassword string `inject:"${mail.smtp.password}"`
}

