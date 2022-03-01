package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateUser, downCreateUser)
}

func upCreateUser(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec("INSERT INTO user (id,created_at,updated_at,first_name,last_name,email,password,phone) VALUES (2,NOW(),NOW(),'Phi','Khanh2','khanh2@gmail.com','pass','phone')")
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func downCreateUser(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DELETE FROM user where id = 2")
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
