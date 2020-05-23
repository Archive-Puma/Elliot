package portscanner

import (
	"errors"
	"strings"

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
	if err := out.DB.CreateTable("PORTSCANNER",
		"`TARGET` INTEGER NOT NULL,`PORT` INT NOT NULL, `PROTOCOL` TEXT NOT NULL, `SERVICE` TEXT NOT NULL, `BANNER` TEXT, PRIMARY KEY(`TARGET`,`PORT`), FOREIGN KEY(`TARGET`) REFERENCES `TARGET`(`ID`)"); err != nil {
		return err
	}
	// Create the statement
	stmt, err := out.DB.Instance.Prepare("INSERT INTO `PORTSCANNER` VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Insert the values
	for _, result := range results {
		chunk := strings.SplitN(result, ",", 4)
		_, _ = stmt.Exec(id, chunk[0], chunk[1], chunk[2], chunk[3])
	}
	return nil
}
