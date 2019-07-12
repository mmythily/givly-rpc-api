package main

import (
	"github.com/rumsrami/givly-rpc-api/http"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// For testing purposes only- remove before prod
const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "givly"
)

func init() {
	// migrationPath, _ := config.GetPath("/assets/sql/migrations")
}

func main() {
	server := http.NewRPCServer()
	server.Run(":8080")
}
