package mysql

import (
	"fmt"
	"github.com/bang-go/kit/db"
	"gorm.io/gorm/schema"
	"testing"
)

func TestConn(t *testing.T) {
	opt := Options{
		Dsn: DsnConfig{User: "test", Passwd: "test", Net: "tcp", Addr: "local:3306", DBName: "test", AllowNativePasswords: true},
		Gorm: db.GormConfig{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	}
	client, err := New(&opt)
	//client, err := New(&Options{
	//	DSN: DSN{Username: "cn5eplay", Password: "cn5E*mysql", Host: "cn5e-dev-mysql-net.mysql.polardb.rds.aliyuncs.com", Port: 3306, Database: "rds5ewin"},
	//})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(client)
	type YwUser struct {
		Uid      int64 `gorm:"primary_key" json:"uid"`
		Username string
	}
	user := &YwUser{}
	client.Sdb.Find(user)
	fmt.Println(user)
}
