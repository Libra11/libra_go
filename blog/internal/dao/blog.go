package dao

import (
	"context"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/database/gorms"
)

type BlogDao struct {
	conn *gorms.GormConn
}

func (b BlogDao) AddTag(ctx context.Context, tag *blogs.Tag) error {
	return b.conn.Session(ctx).Save(tag).Error
}

func (b BlogDao) AddCategory(ctx context.Context, category *blogs.Category) error {
	return b.conn.Session(ctx).Save(category).Error
}

func (b BlogDao) DeleteTag(ctx context.Context, id int64) error {
	return b.conn.Session(ctx).Delete(&blogs.Tag{}, "id = ?", id).Error
}

func (b BlogDao) DeleteCategory(ctx context.Context, id int64) error {
	return b.conn.Session(ctx).Delete(&blogs.Category{}, "id = ?", id).Error
}

func (b BlogDao) GetBlogById(ctx context.Context, id int64) (*blogs.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogDao) GetAllTags(ctx context.Context) ([]*blogs.Tag, error) {
	result := b.conn.Session(ctx).Find(&blogs.Tag{})
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	var tags []*blogs.Tag
	err := result.Scan(&tags)
	if err.Error != nil {
		return nil, err.Error
	}
	return tags, nil
}

func (b BlogDao) GetAllCategory(ctx context.Context) ([]*blogs.Category, error) {
	result := b.conn.Session(ctx).Find(&blogs.Category{})
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	var category []*blogs.Category
	err := result.Scan(&category)
	if err.Error != nil {
		return nil, err.Error
	}
	return category, nil
}

func (b BlogDao) AddBlog(ctx context.Context, blog *blogs.Blog) error {
	return b.conn.Session(ctx).Save(blog).Error
}

func NewBlogDao() *BlogDao {
	return &BlogDao{
		conn: gorms.New(),
	}
}
