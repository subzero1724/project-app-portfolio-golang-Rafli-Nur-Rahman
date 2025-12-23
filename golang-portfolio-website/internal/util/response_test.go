package util

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestWriteErrorAndSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	WriteError(rr, 400, "bad request")

	var res Response
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if res.Success {
		t.Fatalf("expected success false")
	}
	if res.Error != "bad request" {
		t.Fatalf("expected error message, got %s", res.Error)
	}

	rr = httptest.NewRecorder()
	WriteSuccess(rr, 200, map[string]string{"a": "b"}, "ok")
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !res.Success || res.Message != "ok" {
		t.Fatalf("unexpected success payload")
	}
}
