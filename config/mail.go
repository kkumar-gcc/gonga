package config

import (
	"gonga/utils"
)

type MailConfig struct {
	Default string
	Mailers map[string]MailerConfig
	From    MailFromConfig
}

type MailerConfig struct {
	Driver     string
	Host       string
	Port       int
	Encryption string
	Username   string
	Password   string
	LocalDomain string
}

type MailFromConfig struct {
	Address string
	Name    string
}

func LoadMailConfig() *MailConfig {
	return &MailConfig{
		Default: utils.Env("MAIL_MAILER", "smtp"),
		Mailers: map[string]MailerConfig{
			"smtp": {
				Driver:     "smtp",
				Host:       utils.Env("MAIL_HOST", "smtp.mailgun.org"),
				Port:       utils.EnvInt("MAIL_PORT", 587),
				Encryption: utils.Env("MAIL_ENCRYPTION", "tls"),
				Username:   utils.Env("MAIL_USERNAME", ""),
				Password:   utils.Env("MAIL_PASSWORD", ""),
				LocalDomain: utils.Env("MAIL_EHLO_DOMAIN",""),
			},
			"mailgun": {
				Driver: "mailgun",
			},
			"ses": {
				Driver: "ses",
			},
			"postmark":{
				Driver: "postmark",
			},
		},
		From: MailFromConfig{
			Address: utils.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
			Name:    utils.Env("MAIL_FROM_NAME", "Example"),
		},
	}
}
