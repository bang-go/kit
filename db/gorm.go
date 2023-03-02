package db

import (
	"context"
	"database/sql"
	"github.com/bang-go/kit/base/bint"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	Ctx context.Context
	Sdb *gorm.DB
}

type GormConfig gorm.Config

type PoolConfig struct {
	MaxIdleConns    int           //最大空闲连接数,默认2
	MaxOpenConns    int           //最大Open链接数,默认0 不限制
	ConnMaxLifetime time.Duration //链接最大生命周期
	ConnMaxIdleTime time.Duration //链接最大空闲周期
}

// SetPool 设置数据池
func SetPool(sqlDB *sql.DB, config PoolConfig) {
	if !bint.IsEmpty(config.MaxIdleConns) {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if !bint.IsEmpty(config.MaxOpenConns) {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if !bint.IsEmpty(config.ConnMaxLifetime) {
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}
	if !bint.IsEmpty(config.ConnMaxIdleTime) {
		sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	}
}
