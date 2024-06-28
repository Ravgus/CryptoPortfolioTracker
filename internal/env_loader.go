package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	envFile := ".env.local"

	if !isFileExist(envFile) {
		envFile = ".env"
	}

	if !isFileExist(envFile) {
		fmt.Println(".env file not found!")

		os.Exit(1)
	}

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal(err)
	}
}

func isFileExist(envFile string) bool {
	_, err := os.Stat(envFile)

	return !os.IsNotExist(err)
}
