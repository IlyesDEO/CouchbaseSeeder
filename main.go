package main

import (
	"fmt"

	"github.com/kilianp07/CassandraSeeder/pkg/couchbase"
	"github.com/kilianp07/CassandraSeeder/pkg/reader"
)

func main() {
	data, err := reader.Read("./contacts.csv")
	if err != nil {
		panic(err)
	}
	db := couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")
	for _, contact := range data {
		fmt.Println("Migrating Contact: ", contact.Name)
		if err := db.MigrateData(contact); err != nil {
			panic(err)
		}
	}
}
