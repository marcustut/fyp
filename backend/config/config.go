package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/marcustut/fyp/backend/internal/util/environment"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime bool
			Charset   string
			Loc       string
			TLS       bool
		}
	}
	Services struct {
		Slide struct {
			Port int
		}
	}
}

// C is a global config variable.
var C config

// ReadConfigOption is a config option.
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig reads from config file and initialize C.
func ReadConfig(option ReadConfigOption) {
	cwd := filepath.Join(rootDir(), "config")
	viper.AddConfigPath(cwd)

	var readEnvErr error

	// load config based on different environment
	if environment.IsDev() {
		viper.SetConfigName("/config.dev")
		readEnvErr = godotenv.Load(cwd + "/.env.dev")
	} else if environment.IsTest() || (option.AppEnv == environment.Test) {
		viper.SetConfigName("/config.test")
		readEnvErr = godotenv.Load(cwd + "/.env.test")
	} else if environment.IsE2E() || (option.AppEnv == environment.E2E) {
		viper.SetConfigName("/config.e2e")
		readEnvErr = godotenv.Load(cwd + "/.env.e2e")
	} else if environment.IsProd() || (option.AppEnv == environment.Production) {
		viper.SetConfigName("/config.prod")
		readEnvErr = godotenv.Load(cwd + "/.env.prod")
	} else {
		log.Fatal("APP_ENV is not found in environment")
	}
	if readEnvErr != nil {
		log.Fatal("error loading .env file: ", readEnvErr)
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	// replace env in yaml with environment variables
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func getEnvOrPanic(envKey string) string {
	env := os.Getenv(envKey)
	if len(env) == 0 {
		panic("unable to find " + envKey + " in the environment")
	}
	return env
}
