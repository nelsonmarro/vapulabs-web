package handlers

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockEmailSender implements services.EmailSender for testing
type MockEmailSender struct {
	SendContactFormFunc func(name, email, message string) error
}

func (m *MockEmailSender) SendContactForm(name, email, message string) error {
	if m.SendContactFormFunc != nil {
		return m.SendContactFormFunc(name, email, message)
	}
	return nil
}

func TestContactHandler_ServeHTTP(t *testing.T) {
	logger := log.New(bytes.NewBuffer(nil), "", 0)

	tests := []struct {
		name           string
		requestBody    string
		mockSendFunc   func(name, email, message string) error
		expectedStatus int
		expectedInBody string // Substring to check in the response
	}{
		{
			name:           "Successful Submission",
			requestBody:    `{"nombre": "John Doe", "email": "john@example.com", "mensaje": "Hello"}`,
			mockSendFunc:   func(n, e, m string) error { return nil },
			expectedStatus: http.StatusOK,
			expectedInBody: "¡Mensaje Enviado!",
		},
		{
			name:           "Validation Error - Missing Fields",
			requestBody:    `{"nombre": "", "email": "john@example.com", "mensaje": ""}`,
			mockSendFunc:   nil, // Should not be called
			expectedStatus: http.StatusOK,
			expectedInBody: "Por favor completa todos los campos",
		},
		{
			name:           "Email Service Error",
			requestBody:    `{"nombre": "John Doe", "email": "john@example.com", "mensaje": "Hello"}`,
			mockSendFunc:   func(n, e, m string) error { return errors.New("smtp error") },
			expectedStatus: http.StatusOK,
			expectedInBody: "No se pudo enviar el mensaje",
		},
		{
			name:           "Bad JSON Request",
			requestBody:    `{invalid-json}`,
			mockSendFunc:   nil,
			expectedStatus: http.StatusBadRequest,
			expectedInBody: "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockEmailSender{SendContactFormFunc: tt.mockSendFunc}
			handler := NewContactHandler(mockService, logger)

			req := httptest.NewRequest(http.MethodPost, "/contact", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedInBody != "" && !strings.Contains(rr.Body.String(), tt.expectedInBody) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedInBody, rr.Body.String())
			}
		})
	}
}
