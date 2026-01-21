package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestProductHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		productID      string
		expectedTitle  string
		expectedPrice  string
	}{
		{
			name:           "Accountable Holo Product",
			productID:      "holo",
			expectedTitle:  "Accountable Holo",
			expectedPrice:  "$39.99",
		},
		{
			name:           "Unknown Product",
			productID:      "unknown",
			expectedTitle:  "Producto No Encontrado",
			expectedPrice:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewProductHandler()

			req := httptest.NewRequest(http.MethodGet, "/products/"+tt.productID, nil)
			
			// Setup chi context to simulate URL params
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.productID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
			}

			body := rr.Body.String()
			if !strings.Contains(body, tt.expectedTitle) {
				t.Errorf("expected body to contain title %q", tt.expectedTitle)
			}
			if tt.expectedPrice != "" && !strings.Contains(body, tt.expectedPrice) {
				t.Errorf("expected body to contain price %q", tt.expectedPrice)
			}
		})
	}
}
