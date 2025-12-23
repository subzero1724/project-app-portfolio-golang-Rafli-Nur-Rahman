package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"portfolio/internal/dto"
	"portfolio/internal/model"
	"portfolio/internal/service"
)

type mockContactRepoForHandler struct{}

func (m *mockContactRepoForHandler) Create(_ context.Context, c *model.Contact) error {
	c.ID = "ch1"
	return nil
}
func (m *mockContactRepoForHandler) GetAll(_ context.Context) ([]*model.Contact, error) {
	return []*model.Contact{{ID: "ch1", Name: "A", Email: "a@b.com", Subject: "x", Message: "m"}}, nil
}
func (m *mockContactRepoForHandler) GetByID(ctx context.Context, id string) (*model.Contact, error) {
	return nil, nil
}
func (m *mockContactRepoForHandler) UpdateStatus(ctx context.Context, id string, status string) error {
	return nil
}
func (m *mockContactRepoForHandler) Delete(ctx context.Context, id string) error { return nil }

// Test invalid contact request gets 400
func TestSubmitContact_ValidationFails(t *testing.T) {
	svc := service.NewContactService(&mockContactRepoForHandler{})
	h := NewContactHandler(svc, zap.NewExample())

	req := dto.CreateContactRequest{Name: "", Email: "bad", Subject: "s", Message: "m"}
	b, _ := json.Marshal(req)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(b))
	h.SubmitContact(rr, r)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

// Test successful contact creation returns 201
func TestSubmitContact_Success(t *testing.T) {
	svc := service.NewContactService(&mockContactRepoForHandler{})
	h := NewContactHandler(svc, zap.NewExample())

	req := dto.CreateContactRequest{Name: "John", Email: "john@example.com", Subject: "Hello", Message: "msg is long enough"}
	b, _ := json.Marshal(req)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(b))
	h.SubmitContact(rr, r)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}
}

// Test ShowContactForm renders
func TestShowContactForm_Renders(t *testing.T) {
	// try to find templates by walking up like the admin test to be robust
	// this test will fail if templates are not present in the workspace
	svc := service.NewContactService(&mockContactRepoForHandler{})
	h := NewContactHandler(svc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/contact", nil)
	h.ShowContactForm(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
