package htenvier

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"strings"
)

func Env(key string) string {
	return os.Getenv(key)
}

func EnvToLower(key string) string {
	return strings.ToLower(Env(key))
}

func EnvExists(key string) bool {
	return Env(key) != ""
}

func EnvOrDefault(key, defaultValue string) (value string) {
	value = Env(key)
	if value == "" {
		value = defaultValue
	}
	return
}

func GetAppEnv() string {
	return EnvToLower("APP_ENV")
}

/***************** Get mode of the project *********************/

func IsStage() bool {
	return GetAppEnv() == "stage" || GetAppEnv() == "stg"
}

func IsDev() bool {
	return GetAppEnv() == "dev" || GetAppEnv() == "development"
}

func IsProduction() bool {
	return GetAppEnv() == "prod" || GetAppEnv() == "production"
}

func IsDebugMode() bool {
	return EnvToLower("APP_DEBUG") == "true"
}

func EnvBool(key string) bool {
	v := Env(key)
	return v == "1" || EnvToLower(v) == "true"
}

// EnvUint64 parse uint64 number
// also ignore number splitter character(,)
// return error if variable is not numeric
// return nil value and nil error if variable is empty
func EnvUint64(key string) (*uint64, error) {
	str := Env(key)
	if str != "" {
		str = strings.Replace(str, ",", "", -1)
		value, err := strconv.ParseUint(str, 10, 64)
		return &value, err
	}

	return nil, nil
}
