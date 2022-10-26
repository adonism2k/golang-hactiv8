package initializers

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	APPName    string `mapstructure:"APP_NAME"`
	APPUrl     string `mapstructure:"APP_URL"`
	ServerPort string `mapstructure:"APP_PORT"`

	JWTSecret  string `mapstructure:"JWT_SECRET"`
	JWTExpired int    `mapstructure:"JWT_EXPIRED_HOUR"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
