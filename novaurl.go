package novaurl

import (
	"math/rand"
	"time"
	"github.com/MordFustang21/Nova"
)

type Shortener struct {
	urls   map[string]string
	domain string
}

func New() *Shortener {
	s := new(Shortener)
	s.urls = make(map[string]string)

	return s
}

func (s *Shortener) Middleware(r *nova.Request, next func()) {
	// check for short url and redirect
}

// SetDomain sets the short url domain to your domain
func (s *Shortener) SetDomain(d string) {
	s.domain = d
}

// AddURL adds URL to lookup for new requests
func (s *Shortener) AddURL(url string) string {
	shortURL := s.domain + "/" + createRand()
	s.urls[shortURL] = url
	return shortURL
}

// GetURL looks up URL based on shortURL
func (s *Shortener) GetURL(shortURL string) string {
	val, ok := s.urls[shortURL]
	if ok {
		return val
	} else {
		return ""
	}
}

// randChars possible characters to use for url
const randChars = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

// createRand generates unique ID for url
func createRand() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 5)
	for i := range b {
		b[i] = randChars[rand.Intn(len(randChars))]
	}
	return string(b)
}