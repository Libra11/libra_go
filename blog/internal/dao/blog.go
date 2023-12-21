package dao

import (
	"context"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/database/gorms"
)

type BlogDao struct {
	conn *gorms.GormConn
}

func (b BlogDao) AddTag(ctx context.Context, tag *blogs.Tag) (int64, error) {
	err := b.conn.Session(ctx).Save(tag).Error
	if err != nil {
		return 0, err
	}
	return tag.Id, nil
}

func (b BlogDao) AddCategory(ctx context.Context, category *blogs.Category) (int64, error) {
	err := b.conn.Session(ctx).Save(category).Error
	if err != nil {
		return 0, err
	}
	return category.Id, nil
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
	var tags []*blogs.Tag
	result := b.conn.Session(ctx).Find(&tags)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return tags, nil
}

func (b BlogDao) GetAllCategory(ctx context.Context) ([]*blogs.Category, error) {
	var category []*blogs.Category
	result := b.conn.Session(ctx).Find(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
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
