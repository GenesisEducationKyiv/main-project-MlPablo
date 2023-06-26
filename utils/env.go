package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type env interface {
	~string | ~int | ~float64 | ~float32
}

func TryGetEnvDefault[T env](env string, val T) T {
	res, err := getEnv[T](env)
	if err != nil {
		return val
	}

	return res
}

func TryGetEnv[T env](env string) T {
	res, err := getEnv[T](env)
	if err != nil {
		logrus.Fatal(err)
	}

	return res
}

// Generic function for get environment of any type.
func getEnv[T env](env string) (T, error) {
	var stdType T

	envVal := os.Getenv(env)
	if envVal == "" {
		return stdType, fmt.Errorf("cannot get %s environment variable", env)
	}

	switch (any(stdType)).(type) {
	case string:
		return any(envVal).(T), nil
	case int:
		res, err := strconv.Atoi(envVal)
		if err != nil {
			return stdType, fmt.Errorf("cannot get %s environment variable. Error on parse int: %w", env, err)
		}

		return any(res).(T), nil
	case float64:
		res, err := strconv.ParseFloat(envVal, 64)
		if err != nil {
			return stdType, fmt.Errorf("cannot get %s environment variable. Error on parse float64: %w", env, err)
		}

		return any(res).(T), nil
	case float32:
		res, err := strconv.ParseFloat(envVal, 32)
		if err != nil {
			return stdType, fmt.Errorf("cannot get %s environment variable. Error on parse float64: %w", env, err)
		}

		return any(res).(T), nil

	default:
		return stdType, fmt.Errorf("unable to parse type of environment")
	}
}
