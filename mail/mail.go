package mail

// Mail 表示要发送的邮件
type Mail struct {
	FromAddr    string
	FromName    string
	ToAddr      []string
	ToName      string
	Title       string
	ContentType string
	Content     string
}
