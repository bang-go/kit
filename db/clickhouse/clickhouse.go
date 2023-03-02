package clickhouse

import (
	"context"
	"fmt"
	"github.com/bang-go/kit/db"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"net/url"
)

type CkConfig clickhouse.Config

// DSN // clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60
type DsnConfig struct {
	Username   string
	Password   string
	Host1      string
	Port1      int64
	Host2      string
	Port2      int64
	Database   string
	Parameters map[string]string
}

type Options struct {
	Ctx  context.Context
	Dsn  DsnConfig
	Sql  CkConfig
	Gorm db.GormConfig
	Pool db.PoolConfig
}

// New 新建客户端
func New(opt *Options) (c *db.Client, err error) {
	c = &db.Client{}
	gcg := gorm.Config(opt.Gorm)
	opt.setDsn()
	c = &db.Client{}
	c.Sdb, err = gorm.Open(clickhouse.New(clickhouse.Config(opt.Sql)), &gcg)
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

// clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60
func (opt *Options) setDsn() {
	d := opt.Dsn
	dsnStr := fmt.Sprintf("clickhouse://%s:%s@%s:%d,%s:%d/%s", d.Username, d.Password, d.Host1, d.Port1, d.Host2, d.Port2, d.Database)
	if len(d.Parameters) > 0 {
		urlValues := url.Values{}
		for key, value := range d.Parameters {
			urlValues.Set(key, value)
		}
		dsnStr = fmt.Sprintf("%s?%s", dsnStr, urlValues.Encode())
	}
	opt.Sql.DSN = dsnStr
}
