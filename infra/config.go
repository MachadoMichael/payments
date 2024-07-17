package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	DbAddress   string
	DbPassword  string
	DbStores    int
	DbPayments  int
	JwtSecret   string
	LogFilePath string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	envDbStores := os.Getenv("DATABASE_STORES")
	if envDbStores == "" {
		log.Fatal("Error on read database name.")
	}

	dbStores, err := strconv.Atoi(envDbStores)
	if err != nil {
		log.Fatal(err)
		panic("Cannot read envDbName")
	}

	envDbPayments := os.Getenv("DATABASE_PAYMENTS")
	if envDbPayments == "" {
		log.Fatal("Error on read database name.")
	}

	dbPayments, err := strconv.Atoi(envDbStores)
	if err != nil {
		log.Fatal(err)
		panic("Cannot read envDbName")
	}

	Config = &ConfigData{
		DbAddress:   os.Getenv("DATABASE_ADDRESS"),
		DbPassword:  os.Getenv("DATABASE_PASSWORD"),
		DbStores:    dbStores,
		DbPayments:  dbPayments,
		JwtSecret:   os.Getenv("JWT_SECRET"),
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
	}

}
