package emailx

import "fmt"

type SendInfo struct {
	Receiver string
	Message  MessageInfo
}

type SendMultiInfo struct {
	Receivers []string
	Message   MessageInfo
}

type SenderInfo struct {
	Name  string
	Email string
}

func (i SenderInfo) getFrom() string {
	return fmt.Sprintf("%s<%s>", i.Name, i.Email)
}

type MessageType string

const (
	MessageTypeText MessageType = "text"
	MessageTypeHTML MessageType = "html"
)

type MessageInfo struct {
	Type    MessageType
	Subject string
	Content string
}

func (e *Email) Addr() string {
	return fmt.Sprintf("%s:%s", e.host, e.port)
}
