package configs

import "github.com/spf13/viper"

type ServerConfig struct {
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    string `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
	DBName    string `mapstructure:"DB_NAME"`
	JWTsecret string `mapstructure:"JWT_SECRET"`
	// Save for later
	// ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`
}

func LoadServerConfig(path string) (ServerConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	config := ServerConfig{}
	err := viper.ReadInConfig()
	viper.Unmarshal(&config)
	return config, err
}
