package settings

type EmailSetting struct {
	EmailSenderName     string `mapstructure:"email_sender_name"`
	EmailSenderAddress  string `mapstructure:"email_sender_address"`
	EmailSenderPassword string `mapstructure:"email_sender_password"`
}

