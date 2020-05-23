package main

// import elliot "github.com/cosasdepuma/elliot/app"

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./elliot.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS commoncrawl(url VARCHAR2 PRIMARY KEY)"); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO commoncrawl(url) values(?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	stmt.Exec("https://github.com/")
	stmt.Exec("https://github.com/cosasdepuma/")

	//elliot.Entrypoint()
}
