package database

import (
	"errors"

	"gorm.io/gorm"
)

var (
	gormDB          *gorm.DB
	InterError      = errors.New("服务器报错，请稍候再试")
	RequestError    = errors.New("请求参数有误")
	FileFormatError = errors.New("文件格式错误")
)

func Init() error {
	// 取消 MySQL 连接：保持为空实现，避免建立任何数据库连接
	return nil
}

type BaseDao struct {
	Engine *gorm.DB
}

func GetDB() *gorm.DB {
	// return gormDB.Debug()
	return gormDB
}

func (p *BaseDao) GetDB() *gorm.DB {
	return p.Engine
}

func (p *BaseDao) Transaction(db *gorm.DB) {
	p.Engine = db
}
