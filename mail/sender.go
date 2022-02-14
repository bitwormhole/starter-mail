package mail

// Sender ...
// 【inject:"#mail.Sender"】
type Sender interface {
	Send(m *Mail) error
}
