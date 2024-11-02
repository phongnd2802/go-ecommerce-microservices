package user

type Config struct {
	Http  HTTPSetting  `mapstructure:"http"`
	Grpc  GrpcSetting  `mapstructure:"grpc"`
	DB    DBSetting    `mapstructure:"db"`
	Redis RedisSetting `mapstructure:"redis"`
	Email EmailSetting `mapstructure:"email"`
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

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type EmailSetting struct {
	EmailSenderName     string `mapstructure:"email_sender_name"`
	EmailSenderAddress  string `mapstructure:"email_sender_address"`
	EmailSenderPassword string `mapstructure:"email_sender_password"`
}
