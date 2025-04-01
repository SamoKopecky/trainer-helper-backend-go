package crud

import (
	"context"
	"database/sql"
	"testing"
	"trainer-helper/db"

	"github.com/uptrace/bun"
)

func testSetup(t *testing.T) *bun.Tx {
	db := db.GetDbConn("postgresql://trainer_helper:alpharius@localhost/trainer_helper?sslmode=disable", true, "file://../migrations")

	db.RunMigrations()
	tx, err := db.Conn.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		tx.Rollback()
		db.DownMigrations()
		db.Conn.Close()
	})
	return &tx
}
