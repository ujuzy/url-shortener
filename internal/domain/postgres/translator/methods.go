package translator

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"url-shortener/internal/domain/model"
)

func (s *service) Select(ctx context.Context, l *model.Link) (*model.Link, error) {
	link := &model.Link{}
	sqlFormattedUrl := fmt.Sprintf("'%s'", l.Url)
	query := fmt.Sprintf("select * from links where id = %d or url = %s", l.ID, sqlFormattedUrl)
	err := s.db.QueryRowContext(ctx, query).Scan(&link.ID, &link.Url)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (s *service) Insert(ctx context.Context, l *model.Link) error {
	query := fmt.Sprintf("insert into links (url) values ('%s')", l.Url)
	_, err := s.db.Exec(query)
	if err != nil && err.(*pq.Error).Code != "23505" {
		return err
	}
	return nil
}
