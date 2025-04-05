package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"trainer-helper/model"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DbConn struct {
	Conn      *bun.DB
	Migration *migrate.Migrate
	Dsn       string
}

func GetDbConn(dsn string, debug bool, migrationPath string) DbConn {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}
	return DbConn{
		Conn:      db,
		Dsn:       dsn,
		Migration: m,
	}
}

func (d DbConn) RunMigrations() {
	err := d.Migration.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No changes to apply, continuing...")
	} else if err == nil {
		fmt.Println("Applying migrations...")
	} else {
		log.Fatal(err)
	}
}

func (d DbConn) DownMigrations() {
	err := d.Migration.Down()

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No changes to apply, continuing...")
	} else if err == nil {
		fmt.Println("Applying migrations...")
	} else {
		log.Fatal(err)
	}
}

func (d DbConn) SeedDb() {
	timeslots := d.seedTimeslots()
	exerciseIds := d.seedExercises(timeslots[0].Id)
	d.seedWorkSets(exerciseIds)

}

func (d DbConn) seedTimeslots() []model.Timeslot {
	ctx := context.Background()
	const TRAINER_ID = "1"
	var timeslots []model.Timeslot
	timeNow := time.Now()

	for range 7 {
		timeslots = append(timeslots, *model.BuildTimeslot("some name", timeNow, timeNow.Add(1*time.Hour), nil, TRAINER_ID, nil))
		timeNow = timeNow.Add(24 * time.Hour)
	}

	_, err := d.Conn.NewInsert().Model(&timeslots).Exec(ctx)
	if err != nil {
		panic(err)
	}

	return timeslots
}

func (d DbConn) seedExercises(timeslotId int) []int {
	ctx := context.Background()
	exerciseTypes := []model.SetTypeEnum{model.Squat, model.RomanianDeadlift}
	exerciseIds := []int{}

	for i, eType := range exerciseTypes {
		exercise := model.BuildExercise(timeslotId, int(i), "some note", eType)
		_, err := d.Conn.NewInsert().Model(exercise).Exec(ctx)
		if err != nil {
			panic(err)
		}
		exerciseIds = append(exerciseIds, exercise.Id)
	}
	return exerciseIds
}

func (d DbConn) seedWorkSets(exerciseIds []int) {
	ctx := context.Background()
	squatData := []struct {
		reps     int
		inensity string
	}{
		{reps: 4, inensity: "105Kg"},
		{reps: 3, inensity: "105Kg"},
		{reps: 6, inensity: "95kg"},
		{reps: 5, inensity: "95Kg"},
	}
	rdlData := []struct {
		reps     int
		inensity string
	}{
		{reps: 7, inensity: "100Kg"},
		{reps: 7, inensity: "100Kg"},
	}

	for _, squat := range squatData {
		work_set := model.BuildWorkSet(exerciseIds[0], squat.reps, nil, squat.inensity)
		_, err := d.Conn.NewInsert().Model(work_set).Exec(ctx)
		if err != nil {
			panic(err)
		}
	}
	for _, squat := range rdlData {
		work_set := model.BuildWorkSet(exerciseIds[1], squat.reps, nil, squat.inensity)
		_, err := d.Conn.NewInsert().Model(work_set).Exec(ctx)
		if err != nil {
			panic(err)
		}
	}
}
