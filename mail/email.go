package mail

import (
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

type Mail struct {
	Pool *email.Pool //邮件连接池
}

type MailAccount struct {
	Identity string
	Username string
	Password string
	Host     string
}

type MailConfig struct {
	Address   string
	PoolCount int
	Auth      MailAccount
}

//配置邮件发送账户
func (m *Mail) Init(config MailConfig) (err error) {
	m.Pool, err = email.NewPool(
		config.Address,
		config.PoolCount,
		smtp.PlainAuth(config.Auth.Identity, config.Auth.Username, config.Auth.Password, config.Auth.Host),
	)
	return
}

//生成邮件信息
func (m *Mail) CreateMail(subject, from string, receivers []string, Content []byte) (mail *email.Email) {
	e := email.NewEmail()
	e.Subject = subject
	e.From = from
	e.To = receivers
	e.Text = Content
	return e
}

//发送邮件信息
func (m *Mail) SendMail(e *email.Email) (err error) {
	return m.Pool.Send(e, 10*time.Second)
}
