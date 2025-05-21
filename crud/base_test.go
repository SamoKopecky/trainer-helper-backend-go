package crud

import (
	"context"
	"database/sql"
	"testing"
	"trainer-helper/config"
	"trainer-helper/db"
	"trainer-helper/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func testSetupDb(t *testing.T) *bun.Tx {
	config := config.GetConfig()
	db := db.GetDbConn(config.GetDSN(), false, "file://../migrations")
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

func TestInsertMany(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)

	// Arange
	var workSets []model.WorkSet
	for range 2 {
		workSets = append(workSets, *workSetFactory())
	}

	// Act
	if err := crud.InsertMany(&workSets); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Assert
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}

	for i := range workSets {
		workSets[i].Timestamp.SetZeroTimes()
		dbModels[i].Timestamp.SetZeroTimes()
	}

	assert.EqualValues(t, dbModels, workSets, "Work sets should be equal")
}

func TestUndeleteSoftDelete(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)
	var worksets []model.WorkSet
	deleteIds := make([]int, 3)
	for range 3 {
		work_set := workSetFactory()
		crud.Insert(work_set)
		worksets = append(worksets, *work_set)
		deleteIds = append(deleteIds, work_set.Id)
	}
	err := crud.DeleteMany(deleteIds)
	if err != nil {
		t.Fatalf("Failed to delete timeslot: %v", err)
	}

	// Act
	if err := crud.UndeleteMany([]int{worksets[0].Id, worksets[1].Id}); err != nil {
		t.Fatalf("Failed to revert soft delete timeslot: %v", err)
	}

	// Assert
	var dbModels []model.WorkSet
	// Can't use get here because of soft delete
	if err := crud.db.NewSelect().Model(&dbModels).WhereAllWithDeleted().Scan(context.Background()); err != nil {
		t.Fatalf("Failed to get timeslots: %v", err)
	}
	require.Equal(t, 3, len(dbModels), "number of db models is not correct")

	dbModelsMap := make(map[int]model.WorkSet)
	for _, model := range dbModels {
		dbModelsMap[model.Id] = model
	}
	require.Nil(t, dbModelsMap[1].DeletedAt)
	require.Nil(t, dbModelsMap[2].DeletedAt)
	require.NotNil(t, dbModelsMap[3].DeletedAt)
}

func TestDelete(t *testing.T) {
	db := testSetupDb(t)
	crud := NewTimeslot(db)
	var timeslots []model.Timeslot
	for range 2 {
		timeslot := timeslotFactory()
		crud.Insert(timeslot)
		timeslots = append(timeslots, *timeslot)
	}

	// Act
	if err := crud.Delete(timeslots[0].Id); err != nil {
		t.Fatalf("Failed to delete timeslot: %v", err)
	}

	// Assert
	var dbModels []model.Timeslot
	// Can't use get here because of soft delete
	if err := crud.db.NewSelect().Model(&dbModels).WhereAllWithDeleted().Scan(context.Background()); err != nil {
		t.Fatalf("Failed to get timeslots: %v", err)
	}
	require.Equal(t, 2, len(dbModels), "number of db models is not correct")

	dbModelsMap := make(map[int]model.Timeslot)
	for _, model := range dbModels {
		model.SetZeroTimes()
		dbModelsMap[model.Id] = model
	}
	require.NotNil(t, dbModelsMap[1].DeletedAt)
	require.Nil(t, dbModelsMap[0].DeletedAt)
}
