package example

import (
	"j_email/mail"

	"github.com/jordan-wright/email"
)

//QQ邮箱服务器地址
// 接收邮件服务器：pop.qq.com，使用SSL，端口号995
// 接收邮件服务器：imap.qq.com，使用SSL，端口号993
// 发送邮件服务器：smtp.qq.com，使用SSL，端口号465或587

func example() {
	//初始化邮件发送信道
	Ch := make(chan *email.Email, 5)

	//初始化Mail连接
	Mail := mail.Mail{}
	Mail.Init(mail.MailConfig{
		Address:   "smtp.qq.com:587",
		PoolCount: 3,
		Auth: mail.MailAccount{
			Identity: "",
			Username: "XXX@qq.com",
			Password: "123456",
			Host:     "smtp.qq.com",
		},
	})

	for i := 0; i < 5; i++ {
		Ch <- Mail.CreateMail("测试主题", "xxx@qq.com<MailTest>", []string{"yyy@qq.com"}, []byte("test"))
	}

	for {
		m := <-Ch
		Mail.SendMail(m)
	}

}
