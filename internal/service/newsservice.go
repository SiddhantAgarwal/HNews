package service

import (
	"context"

	"github.com/SiddhantAgarwal/HNews/internal/model"
)

// Service : News Service Object
type Service interface {
	GetNews(ctx context.Context, numberOfNews int64) []model.News
}
