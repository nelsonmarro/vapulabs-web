// Package handlers contains HTTP handlers for the web application.
package handlers

import (
	"net/http"

	"github.com/nelsonmarro/vapulabs-web/templates/pages"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	component := pages.Home()
	_ = component.Render(r.Context(), w)
}
