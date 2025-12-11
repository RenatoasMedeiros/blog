package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}
func GetInteger(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valCheckInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valCheckInt

}
