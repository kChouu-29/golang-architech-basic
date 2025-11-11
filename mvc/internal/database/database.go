package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dbname     = os.Getenv("go_api_practice")
	password   = os.Getenv("11042005")
	username   = os.Getenv("root")
	port       = os.Getenv("3306")
	host       = os.Getenv("localhost")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
