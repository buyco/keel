package helper

import (
	"github.com/joho/godotenv"
	"os"
)

// LoadEnvFile loads env var from file
func LoadEnvFile(env string) error {
	path, err := os.Getwd()
	if err != nil {
		return ErrorPrint("Error finding current directory")
	}

	if "" == env {
		env = "development"
	}

	godotenv.Load(path + "/" + ".env." + env + ".local")
	if "test" != env {
		godotenv.Load(path + "/" + ".env.local")
	}
	godotenv.Load(path + "/" + ".env." + env)
	godotenv.Load(path + "/.env")

	return nil
}
