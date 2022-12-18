package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Initialize connection parameters
func InitializeClient(_host string, _port int, _user string, _password string, _dbname string) {
	host = _host
	port = _port
	user = _user
	password = _password
	dbname = _dbname
}

// Returns the connection string
func getConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

// Creates a new client session and returns it (Must be closed after use!)
func openConnection() *sql.DB {
	psqlInfo := getConnectionInfo()
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)

	// Check liveness
	err = db.Ping()
	checkError(err)

	return db
}

// Retrive original URL given the hash
func RetrieveOriginalUrl(hash string) string {
	db := openConnection()
	defer db.Close()

	// Query the database
	rows, err := db.Query("SELECT url FROM urls WHERE hash = $1", hash)
	checkError(err)

	// Read URL
	var originalUrl string
	for rows.Next() {
		err = rows.Scan(&originalUrl)
		checkError(err)
		return originalUrl
	}
	return ""
}

// Store the mapping of the hash to the original URL in Postgres
func StoreUrlMapping(hash string, originalUrl string) {
	db := openConnection()
	defer db.Close()

	// Insert the URL mapping into the database
	_, err := db.Exec("INSERT INTO urls (hash, url) VALUES ($1, $2)", hash, originalUrl)
	checkError(err)
}
