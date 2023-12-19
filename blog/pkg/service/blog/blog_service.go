package blog

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"libra.com/blog/internal/dao"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/repo"
	"libra.com/common/errs"
	"libra.com/grpc/service/blog"
	"libra.com/user/pkg/model"
)

type ServiceBlog struct {
	blog.UnimplementedBlogServiceServer
	cache repo.Cache
	blog  repo.BlogRepo
}

func New() *ServiceBlog {
	return &ServiceBlog{
		cache: dao.Rc,
		blog:  dao.NewBlogDao(),
	}
}

func (l *ServiceBlog) GetAllTags(ctx context.Context, message *blog.GetAllTagsRequest) (*blog.GetAllTagsResponse, error) {
	tags, err := l.blog.GetAllTags(ctx)
	if err != nil {
		zap.L().Error("db GetAllTags error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormGetError)
	}
	tagsMsg := &blog.GetAllTagsResponse{}
	if tags != nil {
		_ = copier.Copy(&tagsMsg.Tags, &tags)
	}
	return tagsMsg, nil
}

func (l *ServiceBlog) GetAllCategory(ctx context.Context, message *blog.GetAllCategoryRequest) (*blog.GetAllCategoryResponse, error) {
	category, err := l.blog.GetAllCategory(ctx)
	if err != nil {
		zap.L().Error("db GetAllCategory error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormGetError)
	}
	categoryMsg := &blog.GetAllCategoryResponse{}
	if category != nil {
		_ = copier.Copy(&categoryMsg.Category, &category)
	}
	return categoryMsg, nil
}

func (l *ServiceBlog) DeleteTag(ctx context.Context, message *blog.DeleteTagRequest) (*blog.DeleteTagResponse, error) {
	err := l.blog.DeleteTag(ctx, message.Id)
	if err != nil {
		zap.L().Error("db DeleteTag error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormDeleteError)
	}
	return &blog.DeleteTagResponse{}, nil
}

func (l *ServiceBlog) DeleteCategory(ctx context.Context, message *blog.DeleteCategoryRequest) (*blog.DeleteCategoryResponse, error) {
	err := l.blog.DeleteCategory(ctx, message.Id)
	if err != nil {
		zap.L().Error("db DeleteCategory error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormDeleteError)
	}
	return &blog.DeleteCategoryResponse{}, nil
}

func (l *ServiceBlog) AddTag(ctx context.Context, message *blog.AddTagRequest) (*blog.AddTagResponse, error) {
	tag := &blogs.Tag{}
	_ = copier.Copy(tag, message.Tag)
	err := l.blog.AddTag(ctx, tag)
	if err != nil {
		zap.L().Error("db AddTag error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &blog.AddTagResponse{}, nil
}

func (l *ServiceBlog) AddCategory(ctx context.Context, message *blog.AddCategoryRequest) (*blog.AddCategoryResponse, error) {
	category := &blogs.Category{}
	_ = copier.Copy(category, message.Category)
	err := l.blog.AddCategory(ctx, category)
	if err != nil {
		zap.L().Error("db AddCategory error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &blog.AddCategoryResponse{}, nil
}

func (l *ServiceBlog) AddBlog(ctx context.Context, message *blog.AddBlogRequest) (*blog.AddBlogResponse, error) {
	bls := &blogs.Blog{}
	_ = copier.Copy(bls, message.Blog)
	err := l.blog.AddBlog(ctx, bls)
	if err != nil {
		zap.L().Error("db AddBlog error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &blog.AddBlogResponse{}, nil
}
