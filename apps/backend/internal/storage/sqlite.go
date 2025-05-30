// apps/backend/internal/storage/sqlite.go
package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var Schema = `
CREATE TABLE IF NOT EXISTS blocks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp INTEGER,
	hash TEXT,
	prev_hash TEXT,
	nonce INTEGER
);

CREATE TABLE IF NOT EXISTS transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	block_hash TEXT,
	sender TEXT,
	recipient TEXT,
	amount REAL
);
`

func InitDB(path string) *sqlx.DB {
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		log.Fatalln("failed to connect to sqlite:", err)
	}
	db.MustExec(Schema)
	return db
}
