package config


import "github.com/spf13/viper"

func LoadConfig(path string, cfgName string, cfgType string, cfgObj interface{}) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("user")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(cfgObj)
	if err != nil {
		return err
	}

	return nil
}