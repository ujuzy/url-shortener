package repo

import (
	"context"
	"url-shortener/internal/domain/model"
)

type LinkService interface {
	Select(ctx context.Context, l *model.Link) (*model.Link, error)
	Insert(ctx context.Context, l *model.Link) error
}
