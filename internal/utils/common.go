package utils

import (
	"os"
)

func GetEnvVar(envVar string, defaultVal string) string {
	varVal := os.Getenv(envVar)

	if varVal == "" {
		return defaultVal
	}

	return varVal
}
