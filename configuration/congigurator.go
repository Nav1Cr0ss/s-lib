package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func NewConfiguration(appCfg Configuration, prefixes ...string) {
	var (
		err    error
		prefix string
	)

	err = loadDotEnv()
	if err != nil {
		log.Fatalf("failed on set envs %s: ", err)
	}

	if len(prefixes) > 0 {
		prefix = prefixes[0]
	}

	if err = envconfig.Process(prefix, appCfg); err != nil {
		_ = envconfig.Usage(prefix, appCfg)
		log.Fatalf("failed on parse configs: %s", err)
	}
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
