package tests

import (
	"Tab_Tracker/db"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func SetupTestDB(t *testing.T) *sql.DB {
	LoadEnv(t)

	connStr := fmt.Sprintf("postgres://postgres.idcuvhjbpxkclqgkxonv:%s@aws-0-us-west-1.pooler.supabase.com:5432/postgres", os.Getenv("TEST_DB_PASSWORD"))
	testDB, err := sql.Open("postgres", connStr)

	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	_, err = testDB.Exec("DROP TABLE IF EXISTS users, user_balances, user_transactions, transaction_requests, interactions")
	if err != nil {
		t.Fatalf("Failed to drop tables: %v", err)
	}

	db.CreateTables(testDB)

	time.Sleep(100 * time.Millisecond)

	return testDB
}

const projectDirName = "Tab Tracker"

// LoadEnv loads env vars from .env
// LoadEnv loads env vars from .env
func LoadEnv(t *testing.T) {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}
	rootPath := re.Find([]byte(cwd))

	envPath := string(rootPath) + `\.env`
	log.Printf("Loading .env file from: %s", envPath)

	err = godotenv.Load(envPath)
	if err != nil {
		t.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}
}

func resetEnv(origEnv []string) func() {
	return func() {
		os.Clearenv()
		for _, e := range origEnv {
			pair := regexp.MustCompile(`^([^=]+)=(.*)$`).FindStringSubmatch(e)
			if len(pair) == 3 {
				os.Setenv(pair[1], pair[2])
			}
		}
	}
}
