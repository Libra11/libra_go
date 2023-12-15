package admin

import (
	"context"
	"libra.com/admin/internal/dao"
	"libra.com/admin/internal/repo"
	"libra.com/grpc/service/admin"
)

type ServiceAdmin struct {
	admin.UnimplementedAdminServiceServer
	cache repo.Cache
}

func New() *ServiceAdmin {
	return &ServiceAdmin{
		cache: dao.Rc,
	}
}

func (*ServiceAdmin) GetUserInfo(context.Context, *admin.UserMessage) (*admin.UserResponse, error) {
	return nil, nil
}
