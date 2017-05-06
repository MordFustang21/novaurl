package novaurl

import (
	"testing"
)

func TestShortener_AddURL(t *testing.T) {
	s := New()
	s.SetDomain("mydomain.com")
	url := s.AddURL("http://www.google.com")
	if len(url) != 18 {
		t.Error("Got invalid url")
	}
}

func TestShortener_GetURL(t *testing.T) {
	s := New()
	s.SetDomain("mydomain.com")
	shortURL := s.AddURL("http://www.google.com")

	fullURL := s.GetURL(shortURL)
	if fullURL != "http://www.google.com" {
		t.Error("Got invalid URL")
	}

	emptyURL := s.GetURL("")
	if emptyURL != "" {
		t.Error("GOt invalid URL")
	}
}

func TestShortener_createRand(t *testing.T) {
	r := createRand()
	if r == "" {
		t.Error("Couldn't get rand string")
	}
}