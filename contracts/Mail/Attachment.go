package mail

type Attachment struct {
	Filename string
	MimeType string
	Content  []byte
	ContentID string
}