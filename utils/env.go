package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func TryGetEnv[T ~string](env string) T {
	res := os.Getenv(env)
	if res == "" {
		logrus.Fatalf("error on getting environment variable: %s", env)
	}

	return T(res)
}

func TryGetEnvDefault[T ~string](env string, defaultVal T) T {
	res := os.Getenv(env)
	if res == "" {
		return defaultVal
	}

	return T(res)
}
