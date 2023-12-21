package blog

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hash"
	"io"
	"libra.com/blog/config"
	"libra.com/blog/internal/dao"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/repo"
	"libra.com/blog/pkg/model"
	"libra.com/common/aliyun"
	"libra.com/common/errs"
	"libra.com/grpc/service/blog"
	"strconv"
	"time"
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
	id, err := l.blog.AddTag(ctx, tag)
	if err != nil {
		zap.L().Error("db AddTag error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &blog.AddTagResponse{
		Id: id,
	}, nil
}

func (l *ServiceBlog) AddCategory(ctx context.Context, message *blog.AddCategoryRequest) (*blog.AddCategoryResponse, error) {
	category := &blogs.Category{}
	_ = copier.Copy(category, message.Category)
	id, err := l.blog.AddCategory(ctx, category)
	if err != nil {
		zap.L().Error("db AddCategory error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &blog.AddCategoryResponse{
		Id: id,
	}, nil
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

func (l *ServiceBlog) GetOssToken(ctx context.Context, message *blog.GetOssTokenRequest) (*blog.GetOssTokenResponse, error) {
	resp, err := aliyun.GenerateStsInfo(config.C.AC.RoleArn, config.C.AC.Regin, config.C.AC.Id, config.C.AC.Secret)
	if err != nil {
		zap.L().Error("GetOssToken GenerateStsInfo error", zap.Error(err))
		return nil, errs.GrpcErr(model.OssError)
	}
	var expireTime int64
	expireTime = 3600
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = strconv.FormatInt(expireEnd, 10)

	//create post policy json
	var c aliyun.ConfigStruct
	uploadDir := "src/"
	host := "https://" + config.C.AC.BucketName + "." + config.C.AC.EndPoint
	c.Expiration = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	c.Conditions = append(c.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(c)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(resp.Credentials.AccessKeySecret))
	_, _ = io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	policyToken := aliyun.PolicyToken{}
	policyToken.AccessKeySecret = resp.Credentials.AccessKeySecret
	policyToken.SecurityToken = resp.Credentials.SecurityToken
	policyToken.AccessId = resp.Credentials.AccessKeyId
	policyToken.Expiration = resp.Credentials.Expiration
	policyToken.Expire = tokenExpire
	policyToken.Signature = string(signedStr)
	policyToken.Dir = uploadDir
	policyToken.Policy = string(debyte)
	policyToken.Host = host
	response := &blog.GetOssTokenResponse{}
	_ = copier.Copy(response, &policyToken)
	return response, nil
}
