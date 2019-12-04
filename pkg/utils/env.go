package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thoas/go-funk"
	"os"
)

// Load envs from file
func LoadEnvFile(envFile string, runningEnv string) error {
	if funk.Contains([]string{"development", "test"}, runningEnv) {
		path, err := os.Getwd()
		if err != nil {
			return ErrorPrint("Error finding current directory")
		}
		err = godotenv.Load(fmt.Sprintf("%s/%s", path, envFile))
		if err != nil {
			return ErrorPrint("Error loading environment file")
		}
	}
	return nil
}
