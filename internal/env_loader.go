package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	fileName := ".env.local"

	if !isFileExist(fileName) {
		fileName = ".env"
	}

	if !isFileExist(fileName) {
		fmt.Println(".env file not found!")

		os.Exit(1)
	}

	if err := godotenv.Load(fileName); err != nil {
		log.Fatal(err)
	}
}

func isFileExist(fileName string) bool {
	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}
