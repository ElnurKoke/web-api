package store

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(name, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, name)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	migrations := []string{"userTable.sql", "formTable.sql"}
	for _, migrationFile := range migrations {
		content, err := ioutil.ReadFile(filepath.Join("migrations", migrationFile))
		if err != nil {
			log.Fatal(err)
		}
		if _, err = db.Exec(string(content)); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Connection to the database was successful")

	return db, nil
}
