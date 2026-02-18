package service

import (
	"crypto/rand"
	"encoding/base64"
	"url-shortener/model"
	"url-shortener/storage"
)

type URLService struct {
	storage *storage.MemoryStorage
}

func NewURLService(storage *storage.MemoryStorage) *URLService {
	return &URLService{storage: storage}
}

func (s *URLService) Shorten(originalURL string) string {
	code := generateCode()

	url := &model.URL{
		Code:        code,
		OriginalURL: originalURL,
		Clicks:      0,
	}

	s.storage.Save(url)
	return code
}

func (s *URLService) GetOriginal(code string) (*model.URL, error) {
	url, err := s.storage.Get(code)
	if err != nil {
		return nil, err
	}

	s.storage.IncrementClicks(code)
	return url, nil
}

func generateCode() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}
