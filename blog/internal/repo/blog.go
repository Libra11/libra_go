package repo

import (
	"context"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/dto"
)

type BlogRepo interface {
	GetBlogById(ctx context.Context, id int64) (*blogs.Blog, error)
	GetAllTags(ctx context.Context) ([]*blogs.Tag, error)
	GetAllCategory(ctx context.Context) ([]*blogs.Category, error)
	DeleteTag(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, id int64) error
	AddTag(ctx context.Context, tag *blogs.Tag) (int64, error)
	AddCategory(ctx context.Context, category *blogs.Category) (int64, error)
	AddBlog(ctx context.Context, blog *blogs.Blog) error
	GetBlogs(ctx context.Context, page, pageSize, categroyId, tagId int64, title string) (*dto.BlogList, error)
}
