package out

import (
	"database/sql"
	"fmt"
	"strings"

	// SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type sDB struct {
	Instance *sql.DB
}

// DB TODO: Doc
var DB = new(sDB)

// CreateTabCreateTable TODO: Doc
func (db *sDB) CreateTable(table string, structure string) error {
	if _, err := db.Instance.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(%s)", table, structure)); err != nil {
		return err
	}
	return nil
}

// GetTargetID TODO: Doc
func (db *sDB) GetTargetID(target string) (int, error) {
	stmt, err := db.Instance.Prepare("SELECT `ID` FROM `TARGET` WHERE `TARGET` = ?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	ids, err := stmt.Query(target)
	if err != nil {
		return -1, err
	}
	defer ids.Close()

	if !ids.Next() {
		return -1, nil
	}
	var id int
	ids.Scan(&id)

	return id, nil
}

// CreateDatabase TODO: Doc
func (db *sDB) CreateDatabase(name string) error {
	if strings.HasSuffix(name, ".db") {
		name = strings.TrimSuffix(name, ".db")
	}
	// Create the database file
	database, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", name))
	if err != nil {
		return err
	}
	db.Instance = database
	if err := db.CreateTable("TARGET",
		"`ID` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `TARGET`	TEXT NOT NULL UNIQUE"); err != nil {
		return err
	}
	return nil
}
