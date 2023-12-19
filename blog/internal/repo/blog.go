package repo

import (
	"context"
	"libra.com/blog/internal/data/blogs"
)

type BlogRepo interface {
	GetBlogById(ctx context.Context, id int64) (*blogs.Blog, error)
	GetAllTags(ctx context.Context) ([]*blogs.Tag, error)
	GetAllCategory(ctx context.Context) ([]*blogs.Category, error)
	DeleteTag(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, id int64) error
	AddTag(ctx context.Context, tag *blogs.Tag) error
	AddCategory(ctx context.Context, category *blogs.Category) error
	AddBlog(ctx context.Context, blog *blogs.Blog) error
}
