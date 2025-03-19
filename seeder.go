package main

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

func seedDb(db bun.DB) {
	personId := seedUsers(db)
	seedTimeslots(db, personId)

}

func seedUsers(db bun.DB) int32 {
	ctx := context.Background()

	user := *model.BuildPerson("Samo Kopecky", "abc@wow.com")
	_, err := db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		panic(err)
	}
	return user.Id
}

func seedTimeslots(db bun.DB, personId int32) {
	ctx := context.Background()
	const TRAINER_ID = 1
	var timeslots []model.Timeslot
	timeNow := time.Now()

	for range 7 {
		timeslots = append(timeslots, *model.BuildTimeslot("some name", timeNow, timeNow.Add(1*time.Hour), TRAINER_ID, &personId))
		timeNow = timeNow.Add(24 * time.Hour)
	}

	_, err := db.NewInsert().Model(&timeslots).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
