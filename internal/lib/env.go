package lib

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort       string             `mapstructure:"SERVER_PORT"`
	JwtSecret        string             `mapstructure:"JWT_SECRET"`
	AppTitle         string             `mapstructure:"APP_TITLE"`
	AppUrl           string             `mapstructure:"APP_URL"`
	ClientUrl        string             `mapstructure:"CLIENT_URL"`
	RajaOngkirAPIKey string             `mapstructure:"RAJA_ONGKIR_API_KEY"`
	RedisUrl         string             `mapstructure:"REDIS_URL"`
	LogOutput        string             `mapstructure:"LOG_OUTPUT"`
	LogLevel         string             `mapstructure:"LOG_LEVEL"`
	Environment      string             `mapstructure:"ENV"`
	Cors             CorsConfig         `mapstructure:",squash"`
	Db               DbConfig           `mapstructure:",squash"`
	GoogleOAuth2     GoogleOAuth2Config `mapstructure:",squash"`
	Midtrans         MidtransConfig     `mapstructure:",squash"`
}

type DbConfig struct {
	Host         string `mapstructure:"DB_HOST"`
	User         string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
	DbName       string `mapstructure:"DB_NAME"`
	Port         int    `mapstructure:"DB_PORT"`
	SslMode      string `mapstructure:"DB_SSL_MODE"`
	MaxIdleTime  int    `mapstructure:"DB_MAX_IDLE_TIME"`
	MaxOpenConns int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"DB_MAX_IDLE_CONNS"`
}

type CorsConfig struct {
	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
}

type MidtransConfig struct {
	MerchantID string `mapstructure:"MIDTRANS_MERCHANT_ID"`
	ClientKey  string `mapstructure:"MIDTRANS_CLIENT_KEY"`
	ServerKey  string `mapstructure:"MIDTRANS_SERVER_KEY"`
}

type GoogleOAuth2Config struct {
	ProjectID    string `mapstructure:"GOOGLE_PROJECT_ID"`
	ClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	ClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	RefreshToken string `mapstructure:"GOOGLE_REFRESH_TOKEN"`
}

func NewEnv() *Env {
	env := &Env{}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	raw := viper.GetString("ALLOWED_ORIGINS")
	env.Cors.AllowedOrigins = strings.Split(raw, ",")
	for i := range env.Cors.AllowedOrigins {
		env.Cors.AllowedOrigins[i] = strings.TrimSpace(env.Cors.AllowedOrigins[i])
	}
	return env
}
