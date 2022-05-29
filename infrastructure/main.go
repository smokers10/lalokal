package main

import (
	"lalokal/infrastructure/database"
	"lalokal/infrastructure/lib"
)

func main() {
	db, err := database.MongoInit().MongoDB()
	if err != nil {
		panic(err)
	}

	lib.CollectionBuilder(db)
}
