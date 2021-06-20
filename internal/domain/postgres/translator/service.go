package translator

import (
	"database/sql"
	"url-shortener/internal/domain/repo"
)

type service struct {
	db *sql.DB
}

func New(db *sql.DB) repo.LinkService {
	return &service{
		db: db,
	}
}
