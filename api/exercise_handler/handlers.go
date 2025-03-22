package exercise_handler

import (
	"log"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"trainer-helper/api"
	"trainer-helper/crud"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	paramId, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}
	params := exerciseGetParams{
		Id: int32(paramId),
	}

	crudExercise := crud.CRUDExercise{Db: cc.Db}
	crudTimeslot := crud.CRUDTimeslot{Db: cc.Db}
	res, err := crudExercise.GetExerciseWorkSets(params.Id)
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice of points so that we can append worksets
	exercisesMap := make(map[int32]*model.FullExercise)

	for _, r := range res {
		val, ok := exercisesMap[r.ExerciseId]
		if !ok {
			val = &model.FullExercise{
				Exercise: r.ToExercise(),
			}
			exercisesMap[r.ExerciseId] = val
		}
		val.WorkSets = append(val.WorkSets, r.ToWorkSet())
		val.WorkSetCount += 1

	}
	exercises := slices.Collect(maps.Values(exercisesMap))
	if len(exercises) == 0 {
		exercises = []*model.FullExercise{}
	}
	apiTimeslot, err := crudTimeslot.GetById(params.Id)
	if err != nil {
		log.Fatal(err)
	}

	return cc.JSON(http.StatusOK, model.FullExercises{
		Timeslot:  apiTimeslot,
		Exercises: exercises,
	})

}
