package chuck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJoke(t *testing.T) {
	// SETUP
	// Create mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method %q; got %q", http.MethodGet, r.Method)
		}
		if r.URL.Path != "/" {
			t.Errorf("Expected path %q; got %q", "/", r.URL.Path)
		}

		// Write mock response
		_, err := w.Write([]byte(`{"value": "This is a test joke."}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
			return
		}
	}))
	defer ts.Close()

	// EXECUTE
	joke, _ := GetJoke(ts.URL)

	// VERIFY
	// didn't want to use any external libs but would normally use https://github.com/stretchr/testify to assert
	if joke != "This is a test joke." {
		t.Errorf("Expected joke %q; got %q", "This is a test joke.", joke)
	}
}
