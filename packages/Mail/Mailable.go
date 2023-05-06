package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gonga/config"
	mail "gonga/contracts/Mail"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Mailable struct {
	From        mail.From
	To          []string
	CC          []string
	BCC         []string
	ReplyTo     []string
	Content     mail.Content
	Attachments []mail.Attachment
	Tags        []string
	Metadata    []string
	Mailer      string
}

func (m *Mailable) BuildContent(content mail.Content) mail.Mail {
	m.Content = content
	return m
}

func (m *Mailable) BuildFrom(from mail.From) mail.Mail {
	m.From = from
	return m
}

func (m *Mailable) BuildTo(to []string) mail.Mail {
	m.To = to
	return m
}

func (m *Mailable) BuildCC(cc []string) mail.Mail {
	m.CC = cc
	return m
}

func (m *Mailable) BuildBCC(bcc []string) mail.Mail {
	m.BCC = bcc
	return m
}

func (m *Mailable) BuildAttach(attachments []mail.Attachment) mail.Mail {
	m.Attachments = attachments
	return m
}

func (m *Mailable) Send(mailer ...string) error {
    if len(mailer) == 0 {
        mailer = []string{"smtp"}
    }
    return SendMail(mailer[0], m)
}



func SendMail(mailer string, mail *Mailable) error {
	// Load mail configuration
	config := config.LoadMailConfig()
	mailConfig := config.Mailers[mailer]

	from := config.From
	to := mail.To
	cc := mail.CC
	bcc := mail.BCC
	replyTo := mail.ReplyTo
	content := mail.Content
	attachments := mail.Attachments

	// Set up mail message
	message := email.NewEmail()
	message.From = from.Address
	message.To = to
	message.Cc = cc
	message.Bcc = bcc
	message.ReplyTo = replyTo
	message.Subject = content.Subject
	message.Text = []byte(content.Text)
	message.HTML = []byte(content.Html)

	for _, attachment := range attachments {
		message.Attach(bytes.NewReader(attachment.Content), attachment.Filename, attachment.MimeType)
	}

	// Send mail
	var auth smtp.Auth
	if mailConfig.Username != "" && mailConfig.Password != "" {
		auth = smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Host)
	}
	addr := fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port)
	if mailConfig.Encryption == "tls" {
		return message.SendWithTLS(addr, auth, &tls.Config{InsecureSkipVerify: true})
	} else {
		return message.Send(addr, auth)
	}
}
