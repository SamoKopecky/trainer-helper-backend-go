package crud

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"trainer-helper/config"
	"trainer-helper/db"
	"trainer-helper/model"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func testSetupDb(t *testing.T) *bun.Tx {
	config := config.GetConfig()
	fmt.Printf("\n\n\n %s \n\n\n", config.GetDSN())
	db := db.GetDbConn(config.GetDSN(), true, "file://../migrations")
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

func TestInsert(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)
	workSet := workSetFactory()

	// Act
	if err := crud.Insert(workSet); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Assert
	var dbModels []model.WorkSet
	if err := db.NewSelect().Model(&dbModels).Scan(context.TODO()); err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Equal(t, 1, len(dbModels))
	workSet.Timestamp.SetZeroTimes()
	dbModels[0].Timestamp.SetZeroTimes()
	assert.Equal(t, *workSet, dbModels[0])
}

func TestGet(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)
	var workSets []model.WorkSet
	for range 3 {
		workSets = append(workSets, *workSetFactory())
	}
	if err := crud.InsertMany(&workSets); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to get work sets: %v", err)
	}

	// Assert
	assert.Equal(t, len(workSets), len(dbModels))
	for i := range workSets {
		workSets[i].Timestamp.SetZeroTimes()
		dbModels[i].Timestamp.SetZeroTimes()
	}
	assert.Equal(t, workSets, dbModels)

}

func TestUpdate(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)
	workSet := workSetFactory()
	crud.Insert(workSet)

	// Act
	workSet.Intensity = "15Kg"
	if err := crud.Update(workSet); err != nil {
		t.Fatalf("Failed to update work sets: %v", err)
	}

	// Assert
	var dbModels []model.WorkSet
	if err := db.NewSelect().Model(&dbModels).Scan(context.TODO()); err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Equal(t, 1, len(dbModels))
	workSet.Timestamp.SetZeroTimes()
	dbModels[0].Timestamp.SetZeroTimes()
	assert.Equal(t, *workSet, dbModels[0])
}
