package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bharabhi01/shorturl-go/internal/repository"
)

type URLService struct {
	repo    *repository.URLRepository
	baseURL string
}

func NewURLService(repo *repository.URLRepository, baseURL string) *URLService {
	return &URLService{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (s *URLService) CreateShortURL(ctx context.Context, longURL string, userID string) (string, error) {
	shortCode := s.generateShortCode(longURL, userID)

	url := &repository.URL{
		ShortCode: shortCode,
		LongURL:   longURL,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().AddDate(1, 0, 0),
	}

	if err := s.repo.Create(url); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", s.baseURL, shortCode), nil
}

func (s *URLService) GetLongURL(ctx context.Context, shortCode string) (string, error) {
	url, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	if url == nil {
		return "", fmt.Errorf("URL not found")
	}

	go func() {
		_ = s.repo.IncrementVisits(shortCode)
	}()

	return url.LongURL, nil
}

func (s *URLService) generateShortCode(longURL string, userID string) string {
	input := fmt.Sprintf("%s:%s:%d", longURL, userID, time.Now().UnixNano())

	hash := sha256.Sum256([]byte(input))

	encoded := base64.URLEncoding.EncodeToString(hash[:6])

	if len(encoded) > 8 {
		return encoded[:8]
	}

	return encoded
}
