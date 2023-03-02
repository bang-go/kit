package mysql

import (
	"context"
	"github.com/bang-go/kit/db"
	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DsnConfig driver.Config //DSN config
type SqlConfig mysql.Config

type Options struct {
	Ctx  context.Context
	Dsn  DsnConfig
	Sql  SqlConfig
	Gorm db.GormConfig
	Pool db.PoolConfig
}

// New 新建客户端
func New(opt *Options) (c *db.Client, err error) {
	c = &db.Client{}
	gcg := gorm.Config(opt.Gorm)
	opt.setDsn() // DSN [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	c.Sdb, err = gorm.Open(mysql.New(mysql.Config(opt.Sql)), &gcg)
	if err != nil {
		return
	}
	c.Ctx = opt.Ctx
	if opt.Ctx == nil {
		c.Ctx = context.Background()
	}
	sqlDb, err := c.Sdb.DB()
	if err != nil {
		return
	}
	db.SetPool(sqlDb, opt.Pool) //设置连接池
	return
}

// 解析DSN CONFIG到字符串
func (opt *Options) setDsn() {
	cf := driver.Config(opt.Dsn)
	opt.Sql.DSN = cf.FormatDSN()
}
