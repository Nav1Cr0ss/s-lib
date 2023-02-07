package cofiguration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	UserName string `required:"true" envconfig:"DB_USERNAME" `
	Port     int    `required:"true" envconfig:"DB_PORT"`
	Host     string `required:"true" envconfig:"DB_HOST"`
	Password string `required:"true" envconfig:"DB_PASSWORD"`
	Name     string `required:"true" envconfig:"DB_NAME"`
}

type Config[C struct{}] struct {
	Db  DBConfig
	App C
}

func NewConfiguration[C struct{}](appCfg C, prefixes ...string) (c *Config[C], err error) {
	var (
		prefix string
	)
	c = &Config[C]{App: appCfg}
	err = loadDotEnv()
	if err != nil {
		return
	}
	if len(prefixes) > 0 {
		prefix = prefixes[0]
	}

	if err = envconfig.Process(prefix, c); err != nil {
		_ = envconfig.Usage(prefix, c)
		log.Fatal("failed on parse configs")
		return
	}

	return
}

func loadDotEnv() error {
	envPath := os.Getenv("ENV_FILE")

	var err error
	if envPath == "" {
		_ = godotenv.Load(".env") // ignore error by default
	} else {
		err = godotenv.Load(envPath) // if path to env file defined, check error
	}

	if err != nil {
		return err
	}

	return nil
}
