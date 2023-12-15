package dao

import (
	"context"
	"gorm.io/gorm"
	"libra.com/user/internal/data/member"
	"libra.com/user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) GetMemberInfo(ctx context.Context, memberId int64) (*member.Member, error) {
	var mem member.Member
	err := m.conn.Session(ctx).Where("id=?", memberId).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &mem, err
}

func (m *MemberDao) FindMember(ctx context.Context, name string, pwd string) (*member.Member, error) {
	var mem member.Member
	err := m.conn.Session(ctx).Where("account=? and password=?", name, pwd).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &mem, err
}

func (m *MemberDao) SaveMember(ctx context.Context, member *member.Member) error {
	return m.conn.Session(ctx).Save(member).Error
}

func (m *MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("account=?", account).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("mobile=?", mobile).Count(&count).Error
	return count > 0, err
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}
