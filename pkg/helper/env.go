package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thoas/go-funk"
	"os"
)

// LoadEnvFile loads env var from file
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
