package user

type Config struct {
	Http HTTPSetting `mapstructure:"http"`
	Grpc GrpcSetting `mapstructure:"grpc"`
	DB   DBSetting   `mapstructure:"db"`
}

type HTTPSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type GrpcSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DBSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
	SslMode  string `mapstructure:"ssl_mode"`
	PoolMax  int    `mapstructure:"pool_max"`
}
