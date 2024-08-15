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
	"context"
	"database/sql"
	"os"
	"time"

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

		CREATE TABLE height_plateaux (
			project_id UUID PRIMARY KEY,
			data JSON,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		);

		-- Magic values
		INSERT INTO projects VALUES ("feedface-cafe-beef-feed-facecafebeef");
	`

	getBuildingLimitsQuery = `
		SELECT data
		FROM building_limits
		WHERE project_id = :project_id
		LIMIT 1
	`

	updateBuildingLimitsQuery = `
		-- Upsert
		INSERT INTO building_limits(project_id,data) VALUES(:project_id, :data)
  		ON CONFLICT(project_id) DO UPDATE SET data=excluded.data;
	`

	getHeightPlateauxQuery = `
		SELECT data
		FROM height_plateaux
		WHERE project_id = :project_id
		LIMIT 1
	`

	updateHeightPlateauxQuery = `
		-- Upsert
		INSERT INTO height_plateaux(project_id,data) VALUES(:project_id, :data)
  		ON CONFLICT(project_id) DO UPDATE SET data=excluded.data;
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

	mnemosyne.db.SetMaxIdleConns(50)
	mnemosyne.db.SetMaxOpenConns(50)
	mnemosyne.db.SetMaxOpenConns(10)

	if err != nil {
		log.Error(err, "failed to open the database")
		panic(err)
	}

	// Load the schema
	if _, err = mnemosyne.db.Exec(scheme); err != nil {
		log.Error(err, "failed to load the scheme")
		panic(err)
	}
	log.V(2).Info("Scheme is loaded successfully")

	return &mnemosyne
}

func (m *Mnemosyne) PingContext(ctx context.Context) error {
	return m.db.PingContext(ctx)
}

// MARK: Building Limits

func (m *Mnemosyne) GetBuildingLimits(projectID string) (string, error) {
	return m.getObject(projectID, getBuildingLimitsQuery)
}

func (m *Mnemosyne) UpdateBuildingLimits(projectID string, data string) (err error) {
	return m.updateObject(projectID, updateBuildingLimitsQuery, data)
}

// MARK: Height plateaux

func (m *Mnemosyne) GetHeightPlateaux(projectID string) (string, error) {
	return m.getObject(projectID, getHeightPlateauxQuery)
}

func (m *Mnemosyne) UpdateHeightPlateaux(projectID string, data string) (err error) {
	return m.updateObject(projectID, updateHeightPlateauxQuery, data)
}

// MARK: Private API

func (m *Mnemosyne) getObject(projectID string, sqlQuery string) (objectData string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		m.log.Error(err, "failed to start the transaction")
		return "", err
	}

	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		m.log.Error(err, "failed to prepare the SQL statement")
	}

	if err = stmt.QueryRowContext(ctx, sql.Named("project_id", projectID)).Scan(&objectData); err != nil {
		if err == sql.ErrNoRows {
			// The project doesn't have any building limits yet
			return "", ErrNotFound
		}

		return "", err
	}

	if err = tx.Commit(); err != nil {
		m.log.Error(err, "failed to commit the transaction")
		return "", err
	}

	return
}

func (m *Mnemosyne) updateObject(projectID string, sqlQuery string, data string) (err error) {
	tx, err := m.db.Begin()
	if err != nil {
		m.log.Error(err, "failed to start the transaction")
		return err
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		m.log.Error(err, "failed to prepare the SQL statement")
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(sql.Named("project_id", projectID), sql.Named("data", data)); err != nil {
		m.log.Error(err, "failed to update the building limits")
		return err
	}

	if err = tx.Commit(); err != nil {
		m.log.Error(err, "failed to commit the transaction")
		return err
	}

	return nil
}

// Home-made destructor. Inspired by Rust.
// Ref: https://rust-unofficial.github.io/patterns/idioms/dtor-finally.html
func (m *Mnemosyne) Drop() {
	m.db.Close()
}
