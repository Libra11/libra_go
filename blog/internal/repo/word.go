package repo

import (
	"context"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/dto"
)

type WordRepo interface {
	AddWord(ctx context.Context, word *blogs.Word) (int64, error)
	GetWords(ctx context.Context, page, pageSize int64) (*dto.WordList, error)
	DeleteWord(ctx context.Context, id int64) error
	GetWordById(ctx context.Context, id int64) (*blogs.Word, error)
	GetWordTranslate(ctx context.Context, word string) (*blogs.Word, error)
}
