package util

import (
	"testing"
)

func TestNewLogger_NotNil(t *testing.T) {
	l := NewLogger()
	if l == nil {
		t.Fatalf("expected logger not nil")
	}
}
