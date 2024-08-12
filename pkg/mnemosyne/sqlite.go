/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

/*
 * Why Mnemosyne? -- https://en.wikipedia.org/wiki/Mnemosyne
 */

package mnemosyne

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-logr/logr"
)

const (
	databasePath = "tmp/mnemosyne.db"

	scheme = `
		CREATE TABLE projects (id UUID PRIMARY KEY);
		CREATE TABLE building_limits (
			project_id UUID PRIMARY KEY,
			data JSON,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		);

		CREATE TABLE height_plateaus (
			project_id UUID PRIMARY KEY,
			data JSON,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		);

		CREATE TABLE split_building_limits (
			project_id UUID PRIMARY KEY,
			data JSON,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		);

		-- Magic values
		INSERT INTO projects VALUES ("feedface-cafe-beef-feed-facecafebeef");
	`

	getBuildingLimitsQuery = `
		SELECT *
		FROM building_limits
		WHERE project_id = ?
		LIMIT 1
	`
)

type Mnemosyne struct {
	log *logr.Logger
	db  *sql.DB
}

func New(log logr.Logger) *Mnemosyne {
	mnemosyne := Mnemosyne{
		log: &log,
	}

	var err error

	// We don't care if the database doesn't exist
	_ = os.Remove(databasePath)

	mnemosyne.db, err = sql.Open("sqlite3", databasePath)

	if err != nil {
		log.Error(err, "failed to open the database")
		panic(err)
	}
	defer mnemosyne.db.Close()

	// Load the schema
	if _, err = mnemosyne.db.Exec(scheme); err != nil {
		log.Error(err, "failed to load the scheme")
		panic(err)
	}
	log.V(2).Info("Scheme is loaded successfully")

	return &mnemosyne
}

func (m *Mnemosyne) GetBuildingLimits() (building_limits any, err error) {
	m.db.Query()
}
