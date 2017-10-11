package sec

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DatabaseInit is used to initialize the database connection and store it a
// newly crafted SecurityContext
func databaseOpen(path string) (*sql.DB, error) {
	if path == "" {
		return nil, fmt.Errorf("DatabaseInit: database path cannot be empty")
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databaseInit(db *sql.DB) error {
	initializationStatement, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS users (
      uid INTEGER PRIMARY KEY AUTOINCREMENT,
      username VARCHAR(64) NOT NULL,
      password VARCHAR(128) NOT NULL,
      key BLOB NULL,
      challenge VARCHAR(128) NULL,
      last_login TIMESTAMP NULL,
      password_expires TIMESTAMP NULL,
      created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS tokens (
      uid INTEGER NOT NULL,
      token VARCHAR(128) NOT NULL,
      expires TIMESTAMP NOT NULL,
      PRIMARY KEY(uid, token),
      FOREIGN KEY (uid) REFERENCES users(uid)
    )`)
	if err != nil {
		return err
	}

	initializationStatement.Exec()

	return nil
}
