package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(path, name string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed!")
	})
	//v.WatchConfig()
	return v
}