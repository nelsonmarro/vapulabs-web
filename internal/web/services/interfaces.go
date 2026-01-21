package services

type EmailSender interface {
	SendContactForm(name, email, message string) error
}
