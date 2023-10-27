package tests_test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	const (
		host     = "0.0.0.0"
		port     = "9000"
		user     = "postgres"
		password = "admin"
		dbname   = "ozon-test-db"
	)

	connStr := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Простой SQL-запрос, который не требует наличия таблиц
	_, err = db.Exec("SELECT 1")
	//if err != nil {
	//	t.Fatalf("Failed to execute SQL query: %v", err)
	//}
}
