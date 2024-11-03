package main

import "github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"

type UserConfig struct {
	Grpc  settings.GrpcSetting     `mapstructure:"grpc"`
	DB    settings.PostgresSetting `mapstructure:"postgres"`
	Redis settings.RedisSetting    `mapstructure:"redis"`
	Email settings.EmailSetting    `mapstructure:"email"`
}
