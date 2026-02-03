package handlers

import (
	"log"
	"net/http"

	"github.com/nelsonmarro/vapulabs-web/internal/web/services"
	"github.com/nelsonmarro/vapulabs-web/templates/components/contact"
	"github.com/starfederation/datastar-go/datastar"
)

type ContactHandler struct {
	emailService *services.EmailService
	errorLog     *log.Logger
}

func NewContactHandler(emailService *services.EmailService, errorLog *log.Logger) *ContactHandler {
	return &ContactHandler{
		emailService: emailService,
		errorLog:     errorLog,
	}
}

type ContactFormSignals struct {
	Nombre  string `json:"nombre"`
	Email   string `json:"email"`
	Mensaje string `json:"mensaje"`
}

func (h *ContactHandler) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	signals := &ContactFormSignals{}
	if err := datastar.ReadSignals(r, signals); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send Email using service
	err := h.emailService.SendContactForm(signals.Nombre, signals.Email, signals.Mensaje)
	if err != nil {
		h.errorLog.Printf("Error sending email: %v", err)
		// We could show an error fragment here, but showing success is safer for UX
		// if we log the failure internally. Or we could send a console error.
		sse := datastar.NewSSE(w, r)
		sse.ConsoleError(err)
	}

	// Respond with Success View
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(contact.ContactSuccess(), datastar.WithSelectorID("contact-form-content"))
}

func (h *ContactHandler) ServeForm(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.ExecuteScript("window.location.reload()")
}
