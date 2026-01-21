package services

import (
	"fmt"

	"github.com/nelsonmarro/vapulabs-web/config"
	"github.com/resend/resend-go/v2"
)

type EmailService struct {
	client *resend.Client
	config *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	client := resend.NewClient(cfg.ResendAPIKey)
	return &EmailService{
		client: client,
		config: cfg,
	}
}

func (s *EmailService) SendContactForm(name, email, message string) error {
	if s.config.ResendAPIKey == "" {
		return fmt.Errorf("resend API key is missing")
	}

	htmlContent := fmt.Sprintf(`
		<p><strong>Nombre:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Mensaje:</strong></p>
		<p>%s</p>
	`, name, email, message)

	params := &resend.SendEmailRequest{
		From:    "NaphSoft <notificaciones@naphsoft.dev>",
		To:      []string{s.config.ContactEmail},
		Subject: fmt.Sprintf("Nuevo Mensaje de Contacto de %s", name),
		Html:    htmlContent,
		ReplyTo: email,
	}

	_, err := s.client.Emails.Send(params)
	return err
}
