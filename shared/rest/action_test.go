package rest

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestActionRestResourceExecute(t *testing.T) {
	tests := []struct {
		name     string
		method   ActionMethod
		wantCall string
	}{
		{"put", ActionMethodPut, "put"},
		{"post", ActionMethodPost, "postWithID"},
		{"patch", ActionMethodPatch, "patch"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var gotCall, gotPath, gotID string
			var gotBody []byte
			record := func(call string) func(InstanaDataObject, string) ([]byte, error) {
				return func(d InstanaDataObject, p string) ([]byte, error) {
					gotCall, gotPath, gotID = call, p, d.GetIDForResourcePath()
					gotBody, _ = json.Marshal(d)
					return []byte("{}"), nil
				}
			}
			client := &mockRestClient{
				putFunc:        record("put"),
				postWithIDFunc: record("postWithID"),
				patchFunc:      record("patch"),
			}

			err := NewActionRestResource(client, "/api/test").Execute(tc.method, "abc", "enable")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotCall != tc.wantCall {
				t.Errorf("expected call %s, got %s", tc.wantCall, gotCall)
			}
			if gotPath != "/api/test" {
				t.Errorf("expected base path /api/test, got %s", gotPath)
			}
			if gotID != "abc/enable" {
				t.Errorf("expected id path abc/enable, got %s", gotID)
			}
			if string(gotBody) != "null" {
				t.Errorf("expected null body, got %s", gotBody)
			}
		})
	}
}

func TestActionRestResourceExecuteReturnsClientError(t *testing.T) {
	expected := errors.New("boom")
	client := &mockRestClient{
		putFunc: func(InstanaDataObject, string) ([]byte, error) { return nil, expected },
	}

	err := NewActionRestResource(client, "/api/test").Execute(ActionMethodPut, "abc", "disable")
	if !errors.Is(err, expected) {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestActionRestResourceExecuteUnsupportedMethod(t *testing.T) {
	err := NewActionRestResource(&mockRestClient{}, "/api/test").Execute(ActionMethod("GET"), "abc", "enable")
	if err == nil {
		t.Fatal("expected error for unsupported method")
	}
}
