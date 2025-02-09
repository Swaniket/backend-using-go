package env

import (
	"os"
	"strconv"
)

func GetEnvVariableAsString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetEnvVariableAsInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val) // Atoi recieves a string & converts it to an integer
	if err != nil {
		return fallback
	}

	return valAsInt
}
