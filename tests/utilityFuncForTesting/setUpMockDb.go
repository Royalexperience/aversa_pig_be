package utilityFuncTest

import (
    "github.com/DATA-DOG/go-sqlmock"
    "os"
	"github.com/joho/godotenv"	
    "fmt"
	"testing"
	"database/sql"
	_ "github.com/lib/pq"

)

func SetUpMockDb(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error){
	godotenv.Load("resources/config-local.env")
	USERSDB_USERNAME := os.Getenv("USERSDB_USERNAME")
	USERSDB_PASSWORD := os.Getenv("USERSDB_PASSWORD")
	USERSDB_NAME := os.Getenv("USERSDB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s", USERSDB_USERNAME, USERSDB_PASSWORD, USERSDB_NAME)
	db, mock, err := sqlmock.NewWithDSN(dsn)
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	return db, mock, nil
}
