package commoncrawler

import (
	"errors"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/out"
)

// Save TODO: Doc
func (plgn Plugin) Save(results []string) error {
	// Check if the target exists in the Database
	id, err := out.DB.GetTargetID(env.Config.Target)
	if err != nil {
		return err
	}
	if id == -1 {
		return errors.New("Error in TARGET Table")
	}
	// Create table if not exists
	if err := out.DB.CreateTable("COMMONCRAWLER",
		"`TARGET` INTEGER NOT NULL,`LINK` TEXT NOT NULL, PRIMARY KEY(`LINK`), FOREIGN KEY(`TARGET`) REFERENCES `TARGET`(`ID`)"); err != nil {
		return err
	}
	// Create the statement
	stmt, err := out.DB.Instance.Prepare("INSERT INTO `COMMONCRAWLER` VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Insert the values
	for _, result := range results {
		_, _ = stmt.Exec(id, result)
	}
	return nil
}
