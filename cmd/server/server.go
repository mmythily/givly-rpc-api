package main

import (
	// "fmt"
	"github.com/rumsrami/givly-rpc-api/pkg/http"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/rumsrami/givly-rpc-api/pkg/storage/psql"
)

// For testing purposes only- remove before prod
const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "givly"
)


func main() {
	// migrationPath, _ := config.GetPath("/assets/sql/migrations")
	// psqlInfo := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable", host, port, user, dbname)
	// db, err := gorm.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	// db.SingularTable(true)
	// db.LogMode(true)

	
	// merchant := &psql.Merchant{}
	// db.Create(merchant)
	server := http.NewRPCServer()
	server.Run(":8080")
}
