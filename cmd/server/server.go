package main

import (
	"fmt"

	"github.com/rumsrami/givly-rpc-api/pkg/adding"
	"github.com/rumsrami/givly-rpc-api/pkg/http"
	"github.com/rumsrami/givly-rpc-api/pkg/listing"
	"github.com/rumsrami/givly-rpc-api/pkg/storage/psql"
)

// For testing purposes only- remove before prod
const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	dbname = "givly"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable", host, port, user, dbname)
	db, err := psql.New(psqlInfo)
	if err != nil {
		panic(err)
	}

	adder := adding.NewService(db)
	lister := listing.NewService(db)

	defer db.DB.Close()
	server := http.NewRPCServer(adder, lister)
	server.Run(":8080")
}
