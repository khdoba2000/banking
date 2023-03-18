package configs

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	conf *Configuration
	once sync.Once
)

// Config loads configuration using atomic pattern
func Config() *Configuration {
	once.Do(func() {
		conf = load()
	})
	return conf
}

// Configuration ...
type Configuration struct {
	HTTPPort    string
	LogLevel    string
	AppName     string
	Environment string

	ServerPort int
	ServerHost string
	ServiceDir string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	CasbinConfigPath           string
	MiddlewareRolesPath        string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
	JWTRefreshKeyExpireDays   int

	CodeToIgnore  string
	PhoneToIgnore string
}

func load() *Configuration {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file: ", err)
	}

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()

	config.Environment = v.GetString("ENVIRONMENT")
	config.HTTPPort = v.GetString("HTTP_PORT")

	config.LogLevel = v.GetString("LOG_LEVEL")

	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")

	config.CasbinConfigPath = v.GetString("CASBIN_CONFIG_PATH")
	config.MiddlewareRolesPath = v.GetString("MIDDLEWARE_ROLES_PATH")
	config.JWTSecretKey = v.GetString("JWT_SECRET_KEY")
	config.JWTRefreshKey = v.GetString("JWT_REFRESH_KEY")

	config.JWTSecretKeyExpireMinutes = v.GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	config.JWTRefreshKeyExpireHours = v.GetInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	config.JWTRefreshKeyExpireDays = v.GetInt("JWT_REFRESH_KEY_EXPIRE_DAYS_COUNT")
	config.CtxTimeout = v.GetInt("CONTEXT_TIMEOUT")

	config.CodeToIgnore = v.GetString("CODE_TO_IGNORE")
	config.PhoneToIgnore = "+998900000000"

	//validate the configuration
	err = config.validate()
	if err != nil {
		log.Fatal("error validating config: ", err)
	}

	return &config
}

func (c *Configuration) validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	return nil
}
