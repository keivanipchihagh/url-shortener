package dotenv

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

// Loads the .env file from the root of the project
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + "url-shortener" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("Problem loading .env file")
		os.Exit(-1)
	}
}

// Reads the value of the given key from the .env file
func ReadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Cannot find '%s' in .env file!", key)
	}
	return value
}

// Reads the value of the given key from the .env file and returns it as integer
func ReadEnvAsInt(key string) int {
	value, err := strconv.Atoi(ReadEnv(key))
	if err != nil {
		log.Fatalf("Cannot parse '%s' as an integer!", key)
		os.Exit(-1)
	}
	return value
}
