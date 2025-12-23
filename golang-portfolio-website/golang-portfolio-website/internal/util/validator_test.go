package util

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

// simple test to verify validator returns errors

func TestValidateStruct(t *testing.T) {
	// When validator is used via ValidateStruct, it should return error for invalid struct
	var v validator.ValidationErrors
	if err := ValidateStruct(&struct {
		Name string `validate:"required"`
	}{}); err == nil {
		t.Fatalf("expected validation error")
	} else {
		if _, ok := err.(validator.ValidationErrors); !ok {
			t.Fatalf("expected ValidationErrors, got %T", err)
		}
	}
	_ = v
}
