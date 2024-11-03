package main

import (
	"fmt"

	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
)

type ProxyConfig struct {
	Http    settings.HTTPSetting `mapstructure:"http"`
	Service Service              `mapstructure:"service"`
}

type Service struct {
	User struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"user"`
}

func (cfg *ProxyConfig) UserAddr() string {
	return fmt.Sprintf("%s:%d", cfg.Service.User.Host, cfg.Service.User.Port)
}

func (cfg *ProxyConfig) HttpAddr() string {
	return fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)
}