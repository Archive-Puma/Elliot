package robots

import (
	"errors"
	"fmt"
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
	if err := out.DB.CreateTable("ROBOTS",
		"`TARGET` INTEGER NOT NULL,`LINK` TEXT NOT NULL, PRIMARY KEY(`TARGET`,`LINK`), FOREIGN KEY(`TARGET`) REFERENCES `TARGET`(`ID`)"); err != nil {
		return err
	}
	// Create the statement
	stmt, err := out.DB.Instance.Prepare("INSERT INTO `ROBOTS` VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Filter the results
	results = filterDisallow(results)
	results = extendedMode(results)
	// Insert the values
	for _, result := range results {
		_, _ = stmt.Exec(id, result)
	}
	return nil
}

func filterDisallow(robots []string) []string {
	disallowed := make([]string, 0)
	for _, robot := range robots {
		if strings.HasPrefix(robot, "Disallow: ") {
			disallowed = append(disallowed, strings.TrimSpace(strings.TrimPrefix(robot, "Disallow: ")))
		}
	}
	return disallowed
}

func extendedMode(robots []string) []string {
	extended := make([]string, 0)
	for _, robot := range robots {
		if strings.HasPrefix(robot, "Allow: ") || strings.HasPrefix(robot, "Disallow: ") || strings.HasPrefix(robot, "/") {
			splits := strings.SplitN(robot, "/", 2)
			url := env.Config.Target
			if !strings.HasSuffix(url, "/") {
				url = fmt.Sprintf("%s%s", url, "/")
			}
			robot = fmt.Sprintf("%s%s%s", splits[0], url, splits[1])
		}
		extended = append(extended, robot)
	}
	return extended
}
