package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	fmt.Println("init 1_create_users_table")
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id                  SERIAL PRIMARY KEY,

			created_at          TIMESTAMP NOT NULL,
			updated_at          TIMESTAMP NOT NULL,
			deleted_at			TIMESTAMP,

			user_id				UUID NOT NULL,
			
			name				VARCHAR,
			email				VARCHAR NOT NULL,
			phone_number		VARCHAR NOT NULL
		);`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}
