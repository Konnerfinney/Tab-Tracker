package db

import (
    "database/sql"
    "fmt"
    "log"
	"os"
    _ "github.com/lib/pq"

    "github.com/joho/godotenv"
    
)

func ConnectDB() *sql.DB {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    connStr := fmt.Sprintf("postgres://%s:%s@bubble.db.elephantsql.com/%s",os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_USERNAME"))
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
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`
    _, err := db.Exec(query)
    if err != nil {
        log.Fatalf("Could not create table: %v", err)
    }
    fmt.Println("Table created successfully!")
}

func InsertTestRecords(db *sql.DB) {
    users := []struct {
        Name  string
        Email string
    }{
        {"Alice", "alice@example.com"},
        {"Bob", "bob@example.com"},
        {"Charlie", "charlie@example.com"},
    }

    for _, user := range users {
        _, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
        if err != nil {
            log.Fatalf("Could not insert test record: %v", err)
        }
    }
    fmt.Println("Test records inserted successfully!")
}

