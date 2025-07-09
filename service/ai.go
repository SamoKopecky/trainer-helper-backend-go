package service

import (
	"fmt"
	"trainer-helper/fetcher"
	"trainer-helper/store"
)

type AI struct {
	Fetcher           fetcher.AI
	ExerciseTypeStore store.ExerciseType
}

func (ai AI) GenerateWeekDay(trainerId string, rawString string) error {
	exercises, err := ai.ExerciseTypeStore.GetByUserId(trainerId)
	if err != nil {
		return err
	}
	exericseNames := make([]string, len(exercises))

	for i := range exercises {
		exericseNames[i] = exercises[i].Name
	}

	resultJson, err := ai.Fetcher.RawStringToJson(exericseNames, rawString)
	if err != nil {
		return err
	}
	fmt.Println("%s", resultJson)
	return nil
}
