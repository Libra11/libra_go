package dao

import (
	"context"
	"github.com/jinzhu/copier"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/database/gorms"
	"libra.com/blog/internal/dto"
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

func (b BlogDao) GetBlogs(ctx context.Context, page, pageSize, categoryId, tagId int64, title string) (*dto.BlogList, error) {
	var res = &dto.BlogList{
		Blogs: make([]*dto.BlogInfo, 0),
		Total: 0,
	}
	var blogs []*blogs.Blog
	query := b.conn.Session(ctx).Preload("Category").Preload("Tags")

	if categoryId != 0 {
		query = query.Joins("JOIN blog_categories ON blog_categories.blog_id = blogs.id").
			Where("blog_categories.category_id = ?", categoryId)
	}

	if tagId != 0 {
		query = query.Joins("JOIN blog_tags ON blog_tags.blog_id = blogs.id").
			Where("blog_tags.tag_id = ?", tagId)
	}

	if title != "" {
		query = query.Where("blogs.title LIKE ?", "%"+title+"%")
	}

	// 获取总数
	err := query.Model(&blogs).Count(&res.Total).Error
	if err != nil {
		return nil, err
	}

	// 分页查询
	err = query.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&blogs).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&res.Blogs, &blogs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewBlogDao() *BlogDao {
	return &BlogDao{
		conn: gorms.New(),
	}
}
