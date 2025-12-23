package service

import (
	"context"
	"testing"
	"time"

	"portfolio/internal/dto"
	"portfolio/internal/model"
)

type mockContactRepo struct{}

func (m *mockContactRepo) Create(ctx context.Context, c *model.Contact) error {
	c.ID = "c1"
	c.CreatedAt = time.Now()
	return nil
}
func (m *mockContactRepo) GetAll(ctx context.Context) ([]*model.Contact, error) {
	return []*model.Contact{{ID: "c1", Name: "A", Email: "a@b.com", Subject: "Hi", Message: "msg", CreatedAt: time.Now()}}, nil
}
func (m *mockContactRepo) GetByID(ctx context.Context, id string) (*model.Contact, error) {
	return nil, nil
}
func (m *mockContactRepo) UpdateStatus(ctx context.Context, id string, status string) error {
	return nil
}
func (m *mockContactRepo) Delete(ctx context.Context, id string) error { return nil }

func TestCreateContactAndGetAll(t *testing.T) {
	repo := &mockContactRepo{}
	svc := NewContactService(repo)

	req := &dto.CreateContactRequest{Name: "John", Email: "john@example.com", Subject: "Hello", Message: "This is a message"}
	resp, err := svc.CreateContact(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID == "" {
		t.Fatalf("expected id set")
	}

	all, err := svc.GetAllContacts(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(all) == 0 {
		t.Fatalf("expected contacts")
	}
}
