package infrastructure

import "github.com/jmoiron/sqlx"

func AutoMigrate(db *sqlx.DB) {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );`
	db.MustExec(schema)
}
