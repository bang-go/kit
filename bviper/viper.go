package bviper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Options 可选项
type Options struct {
	ConfigType  string   //配置文件的格式
	ConfigPaths []string // 查找配置文件的路径
	ConfigNames []string //配置文件名称(无扩展名)
}

// Configure 初始化配置//TODO:配置中心etc等
func Configure(opt *Options) (err error) {
	viper.SetConfigType(opt.ConfigType) // 查找配置文件的格式
	for _, value := range opt.ConfigPaths {
		viper.AddConfigPath(value) // 查找配置文件的路径
	}
	for _, value := range opt.ConfigNames {
		viper.SetConfigName(value)                   //配置文件名称(无扩展名)
		if err = viper.MergeInConfig(); err != nil { // 处理错误
			return err
		}
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig() //todo 多文件模式 热加载不会生效
	return
}
