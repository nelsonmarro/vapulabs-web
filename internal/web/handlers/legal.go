package handlers

import (
	"net/http"
	"github.com/nelsonmarro/vapulabs-web/templates/pages"
)

type LegalHandler struct{}

func NewLegalHandler() *LegalHandler {
	return &LegalHandler{}
}

func (h *LegalHandler) ServePrivacy(w http.ResponseWriter, r *http.Request) {
	component := pages.LegalPage("Política de Privacidad", pages.PrivacyContent())
	component.Render(r.Context(), w)
}

func (h *LegalHandler) ServeTerms(w http.ResponseWriter, r *http.Request) {
	component := pages.LegalPage("Aviso Legal", pages.TermsContent())
	component.Render(r.Context(), w)
}
