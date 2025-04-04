package crud

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"
	"trainer-helper/db"

	"github.com/uptrace/bun"
)

func randomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func testSetup(t *testing.T) *bun.Tx {
	db := db.GetDbConn("postgresql://trainer_helper:alpharius@localhost/trainer_helper?sslmode=disable", true, "file://../migrations")
	db.DownMigrations()

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
