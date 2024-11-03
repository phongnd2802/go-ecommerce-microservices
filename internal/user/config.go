package user

import "github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"

type Config struct {
	Http  HTTPSetting              `mapstructure:"http"`
	Grpc  GrpcSetting              `mapstructure:"grpc"`
	DB    settings.PostgresSetting `mapstructure:"postgres"`
	Redis settings.RedisSetting    `mapstructure:"redis"`
	Email settings.EmailSetting    `mapstructure:"email"`
}

type HTTPSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type GrpcSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
