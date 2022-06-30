package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	fmt.Println("init 2_create_cars_table")
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Println("creating table cars...")
			_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cars (
			id				SERIAL PRIMARY KEY,

			rent_at         TIMESTAMP,

			car_id			UUID NOT NULL,
			renter_id		UUID,
			
			model				VARCHAR NOT NULL,
			brand				VARCHAR NOT NULL,
			availability		BOOLEAN
		);`)
			if err != nil {
				return err
			}
			_, err = db.Exec(`
			INSERT INTO cars (car_id, brand, model, availability) VALUES
    ('7c61be06-f8ea-424f-b8b6-4a8f54da4b44', 'Mercedes', 'A-Class', true),
	('10c39cfd-fa51-4950-923a-07a52c2536de', 'Mercedes', 'B-Class', true),
	('3aab2511-b1c0-4619-8e5b-ab83d71a358d', 'Mercedes', 'C-Class', true),
	('9b64bca4-b20c-45d8-9e53-61509a3795cc', 'Mercedes', 'E-Class', true),
	('be61c465-002e-4f1b-a8e5-46ac3be38eec', 'Mercedes', 'G-Class', true),
	('c763e609-7b53-4498-b3d8-9254d3d3dfc8', 'Mercedes', 'V-Class', true),
	('01a7400f-b9fd-432f-b57b-b607ee28efa4', 'Mercedes', 'S-Class', true),
	('f932cd78-33ce-45a8-9705-37808a1c2aaa', 'Mercedes', 'T-Class', true),
	('f9fc9e7d-2d28-41ac-b34e-fee3979a17b7', 'Mercedes', 'AMG', true),
	('dd337557-afe5-4db0-98e4-2fc8752bce7a', 'Mercedes', 'CLA', true);
	`)
			return err
		}, func(db migrations.DB) error {
			fmt.Println("dropping table cars...")
			_, err := db.Exec(`DROP TABLE cars`)
			return err
		})
}
