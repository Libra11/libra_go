package blog

import (
	"libra.com/blog/internal/dao"
	"libra.com/blog/internal/repo"
	"libra.com/grpc/service/blog"
)

type ServiceBlog struct {
	blog.UnimplementedAdminServiceServer
	cache repo.Cache
}

func New() *ServiceBlog {
	return &ServiceBlog{
		cache: dao.Rc,
	}
}
