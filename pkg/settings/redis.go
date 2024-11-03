package settings

import "fmt"


type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}


func (rt *RedisSetting) Addr() string {
	return fmt.Sprintf("%s:%d", rt.Host, rt.Port)
}