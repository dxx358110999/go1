package email_provider

import (
	"crypto/tls"
	"dxxproject/config_prepare/app_config"
	"fmt"
	"github.com/samber/do/v2"
	"net/smtp"
)

type SmtpIMPL struct {
	SmtpServer  string // QQ邮箱SMTP服务器地址
	SmtpPort    int    // SMTP端口：465（SSL）或 587（TLS）
	SenderEmail string // 发件人邮箱（你的QQ邮箱）
	AuthCode    string // QQ邮箱授权码（不是登录密码）
}

func (rec *SmtpIMPL) SendEmail(toEmail string, subject string, body string) error {
	/*
		发送邮件到指定邮箱
	*/

	// 1. 配置SMTP认证信息
	auth := smtp.PlainAuth(
		"",
		rec.SenderEmail,
		rec.AuthCode,
		rec.SmtpServer,
	)

	// 3. 拼接SMTP地址（服务器+端口）
	addr := fmt.Sprintf("%s:%d", rec.SmtpServer, rec.SmtpPort)

	// 4. 配置TLS（QQ邮箱要求加密连接）
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（生产环境建议关闭）
		ServerName:         rec.SmtpServer,
	}

	// 5. 建立连接并发送邮件
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("tls连接失败: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, rec.SmtpServer)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %v", err)
	}
	defer client.Close()

	// 认证
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP认证失败: %v", err)
	}

	// 设置发件人
	if err := client.Mail(rec.SenderEmail); err != nil {
		return fmt.Errorf("设置发件人失败: %v", err)
	}

	// 设置收件人
	if err := client.Rcpt(toEmail); err != nil {
		return fmt.Errorf("设置收件人失败: %v", err)
	}

	// 发送邮件内容
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("获取数据写入器失败: %v", err)
	}
	defer w.Close()

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=UTF-8;\n" // 邮件头部（必须按SMTP格式构建）
	msg := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n%s\r\n\r\n%s",
		toEmail,
		rec.SenderEmail,
		subject,
		mime,
		body,
	))

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("写入邮件内容失败: %v", err)
	}

	return nil
}

func NewSmtpIMPL(injector do.Injector) (sender *SmtpIMPL, err error) {
	cfg := do.MustInvoke[*app_config.AppConfig](injector).EmailSmtp
	sender = &SmtpIMPL{
		SmtpServer:  cfg.SmtpHost,
		SmtpPort:    cfg.SmtpPort,
		SenderEmail: cfg.SenderEmail,
		AuthCode:    cfg.AuthCode,
	}
	return
}

var _ EmailProviderIF = new(SmtpIMPL)
