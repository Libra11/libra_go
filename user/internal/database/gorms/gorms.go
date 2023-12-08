package gorms

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"libra.com/user/config"
)

var _db *gorm.DB

func InitDB() {
	//配置MySQL连接参数
	username := config.C.MC.Username //账号
	password := config.C.MC.Password //密码
	host := config.C.MC.Host         //数据库地址，可以是Ip或者域名
	port := config.C.MC.Port         //数据库端口
	Dbname := config.C.MC.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}

func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	db *gorm.DB
}

func New() *GormConn {
	return &GormConn{db: GetDB()}
}
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}
