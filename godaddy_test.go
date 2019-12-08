package godaddy

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("API", "SECRET", false)
	if err != nil {
		t.Error("New client factory returned error")
	}

	if client.Key != "API" {
		t.Errorf("Client.Key incorrect, expected %s, want %s", "API", client.Key)
	}

	if client.Secret != "SECRET" {
		t.Errorf("Client.Key incorrect, expected %s, want %s", "SECRET", client.Secret)
	}
}

func TestNewClientWithEnv(t *testing.T) {
	os.Setenv("GODADDY_API_KEY", "API")
	os.Setenv("GODADDY_API_SECRET", "SECRET")

	client, err := NewClient("", "", false)
	if err != nil {
		t.Error("New client factory returned error")
	}

	if client.Key != "API" {
		t.Errorf("Client.Key incorrect, expected %s, want %s", "API", client.Key)
	}

	if client.Secret != "SECRET" {
		t.Errorf("Client.Key incorrect, expected %s, want %s", "SECRET", client.Secret)
	}

	os.Unsetenv("GODADDY_API_KEY")
	os.Unsetenv("GODADDY_API_SECRET")
}

func TestNewClientFailure(t *testing.T) {
	_, err := NewClient("", "", false)
	if err == nil {
		t.Error("Expected failure, returned success")
	}
}

func TestGetURL(t *testing.T) {
	c := new(Client)
	var url string

	c.OTE = true
	url = c.GetURL()
	if url != "https://api.ote-godaddy.com" {
		t.Errorf("GetURL failure, expected: %s, got: %s", "https://api.ote-godaddy.com", url)
	}

	c.OTE = false
	url = c.GetURL()
	if url != "https://api.godaddy.com" {
		t.Errorf("GetURL failure, expected: %s, got: %s", "https://api.godaddy.com", url)
	}

	os.Setenv("GODADDY_API_URL", "https://example.com")
	url = c.GetURL()
	if url != "https://example.com" {
		t.Errorf("GetURL failure, expected: %s, got: %s", "https://example.com", url)
	}
	os.Unsetenv("GODADDY_API_URL")
}

func TestGenerateError(t *testing.T) {
	e := new(Error)
	e.Message = "Test"

	if e.Error() != "Test" {
		t.Errorf("Error failure, expected: %s, got: %s", "Test", e.Error())
	}
}
