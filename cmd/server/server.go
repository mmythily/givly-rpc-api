package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rumsrami/givly-rpc-api/http"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "givly"
)

func init() {
	// migrationPath, _ := config.GetPath("/assets/sql/migrations")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec(`
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func main() {
	server := http.NewRPCServer()
	server.Run(":8080")
}
