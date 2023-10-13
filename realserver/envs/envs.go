package envs

import (
	"os"
	"strconv"
)

func GetString(name string) string {
	return os.Getenv(name)
}

func GetBool(name string) (bool, error) {
	env := os.Getenv(name)
	return strconv.ParseBool(env)
}
