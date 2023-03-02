package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestViper(t *testing.T) {
	err := InitConfig(&Options{ConfigPaths: []string{"./", "./t"}, ConfigType: "yaml", ConfigNames: []string{"t1", "t2", "t3"}})
	fmt.Println(err)
	fmt.Println(viper.GetString("say"))
	fmt.Println(viper.GetString("send"))
	time.Sleep(60 * time.Second)
	fmt.Println(viper.GetString("send"))

}
