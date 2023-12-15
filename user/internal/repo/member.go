package repo

import (
	"golang.org/x/net/context"
	"libra.com/user/internal/data/member"
)

type MemberRepo interface {
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(ctx context.Context, member *member.Member) error
	FindMember(ctx context.Context, name string, pwd string) (*member.Member, error)
	GetMemberInfo(ctx context.Context, memberId int64) (*member.Member, error)
}
