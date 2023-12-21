package model

import "libra.com/common/errs"

var (
	GormGetError    = errs.NewError(1001, "数据库读取失败")
	GormSetError    = errs.NewError(1002, "数据库存储失败")
	GormDeleteError = errs.NewError(1003, "数据库删除失败")

	OssError = errs.NewError(2000, "oss错误")
)
