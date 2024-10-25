package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/sijms/go-ora"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <username> <password> <connection_string>", os.Args[0])
	}

	username := os.Args[1]
	password := os.Args[2]
	connString := os.Args[3]

	dsn := fmt.Sprintf("oracle://%s:%s@%s", username, password, connString)
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	defer db.Close()

	var result int
	err = db.QueryRow("SELECT 1 FROM dual").Scan(&result)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	fmt.Printf("Query result: %d\n", result)
}
