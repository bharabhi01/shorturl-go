package repository

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type URLRepository struct {
	db    *Database
	cache *Cache
}

func NewURLRepository(db *Database, cache *Cache) *URLRepository {
	return &URLRepository{
		db:    db,
		cache: cache,
	}
}

func (r *URLRepository) Create(url *URL) error {
	return r.db.DB.Create(url).Error
}

func (r *URLRepository) FindByShortCode(ctx context.Context, shortCode string) (*URL, error) {
	longURL, err := r.cache.Get(ctx, shortCode)
	if err == nil {
		return &URL{
			ShortCode: shortCode,
			LongURL:   longURL,
		}, nil
	} else if err != redis.Nil {
		return nil, err
	}

	var url URL
	if err := r.db.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	r.cache.Set(ctx, shortCode, url.LongURL, 24+time.Hour)

	return &url, nil
}

func (r *URLRepository) IncrementVisits(shortCode string) error {
	return r.db.DB.Model(&URL{}).Where("short_code = ?", shortCode).UpdateColumn("visits", gorm.Expr("visits + 1", 1)).Error
}
