package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetDatabaseFromEnv() *sql.DB {
	godotenv.Load()

	DSN := os.Getenv("POSTGRES_DSN") // Data Source Name
	fmt.Println(DSN)

	// Setup DB
	sqlDB, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB
}


func 