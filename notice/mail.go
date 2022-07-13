package notice

import (
	"errors"
	"net/smtp"
)

type Email struct {
	Sender string 
	Recipient []string
	Title string 
	Body string
}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) SetSender(sender string) *Email {
	e.Sender = sender
	return e
}

func (e *Email) To(recipients ...string) *Email {
	e.Recipient = recipients
	return e
}

func (e *Email) Title(title string) *Email {
	e.Title = title
	return e
}

func (e *Email) Body(body string) *Email {
	e.Body = body
	return e
}

func (e *Email) Send() error {
	if len(e.Recipient) == 0 {
		return errors.New("no-recipient-provided")
	} else if e.Title == "" {
		return errors.New("no-title-provided")
	} else if e.Body == "" {
		return errors.New("no-context-provided")		
	}




}