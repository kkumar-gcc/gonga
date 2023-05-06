package mail

type Mail interface {
	BuildContent(content Content) Mail
	BuildFrom(from From) Mail
	BuildTo(to []string) Mail
	BuildCC(cc []string) Mail
	BuildBCC(bcc []string) Mail
	BuildAttach(attachments []Attachment) Mail
	Send(mailer ...string) error
}

type Content struct {
	Subject string
	Html    string
	Text    string
}

type From struct {
	Address string
	Name    string
}