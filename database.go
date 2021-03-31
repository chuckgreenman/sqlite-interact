package sqliteutils

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	name string
	path string
}

func CreateDatabase(path string, inMemory bool) *Database {
	fmt.Println(path)

	if inMemory == false {
		db, err := sql.Open("sqlite3", "./foo.db")

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		return db
	} else {
		return 
	}

	return db
}

func DestroyDatabase(path string) {
	fmt.Println("Destroying a database")
}

func ConnectToDatabase(path string) {
	fmt.Println("Connecting to a database")
}

func createFileDatabase(path string) err {

}
