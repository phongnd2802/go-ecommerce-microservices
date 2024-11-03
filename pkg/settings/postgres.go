package settings

import "fmt"

type PostgresSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
	SslMode  string `mapstructure:"ssl_mode"`
}

func (p *PostgresSetting) Addr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
	 		p.Username, p.Password, p.Host, p.Port, p.DbName, p.SslMode)
}