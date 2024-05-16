package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("postgres://postgres.smtguquakjhntnisovgi:%s@aws-0-us-west-1.pooler.supabase.com:5432/postgres", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() // This ensures that the connection is established
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")
	return db
}

func CreateTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        provider_id TEXT NOT NULL,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        provider TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
	fmt.Println("Table created successfully!")
}
