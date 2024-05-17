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

func CreateTables(db *sql.DB) {
	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        user_id SERIAL PRIMARY KEY,
        provider_id TEXT UNIQUE NOT NULL,
        username TEXT NOT NULL,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        provider TEXT NOT NULL,
        is_proxy_user BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	userBalanceTable := `
    CREATE TABLE IF NOT EXISTS user_balances (
        balance_id SERIAL PRIMARY KEY,
        user1_id INTEGER NOT NULL,
        user2_id INTEGER NOT NULL,
        net_balance DECIMAL NOT NULL,
        FOREIGN KEY (user1_id) REFERENCES users(user_id),
        FOREIGN KEY (user2_id) REFERENCES users(user_id),
        UNIQUE (user1_id, user2_id)
    );`

	userTransactionTable := `
    CREATE TABLE IF NOT EXISTS user_transactions (
        transaction_id SERIAL PRIMARY KEY,
        creditor_id INTEGER NOT NULL,
        debtor_id INTEGER NOT NULL,
        amount DECIMAL NOT NULL,
        description TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (creditor_id) REFERENCES users(user_id),
        FOREIGN KEY (debtor_id) REFERENCES users(user_id)
    );`

	transactionRequestTable := `
    CREATE TABLE IF NOT EXISTS transaction_requests (
        request_id SERIAL PRIMARY KEY,
        transaction_id INTEGER NOT NULL,
        requester_id INTEGER NOT NULL,
        requestee_id INTEGER NOT NULL,
        request_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        status TEXT NOT NULL,
        response_date TIMESTAMP,
        FOREIGN KEY (transaction_id) REFERENCES user_transactions(transaction_id),
        FOREIGN KEY (requester_id) REFERENCES users(user_id),
        FOREIGN KEY (requestee_id) REFERENCES users(user_id)
    );`

	interactionTable := `
    CREATE TABLE IF NOT EXISTS interactions (
        interaction_id SERIAL PRIMARY KEY,
        interaction_starter INTEGER NOT NULL,
        interaction_receiver INTEGER NOT NULL,
        interaction_count INTEGER DEFAULT 0,
        last_interaction_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (interaction_starter) REFERENCES users(user_id),
        FOREIGN KEY (interaction_receiver) REFERENCES users(user_id)
    );`

	tables := []string{
		userTable,
		userBalanceTable,
		userTransactionTable,
		transactionRequestTable,
		interactionTable,
	}

	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatalf("Could not create table: %v", err)
		}
	}
	fmt.Println("Tables created successfully!")
}
