package exercise_type_handler

import (
	"errors"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*schemas.DbContext)
	var exerciseTypes []model.ExerciseType

	params, err := api.BindParams[exericseTypeGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	if params.UserId == "" {
		return cc.BadRequest(errors.New("User id parameter missing"))
	}

	exerciseTypes, err = cc.ExerciseTypeCrud.GetByUserId(params.UserId)
	if err != nil {
		return err
	}

	if len(exerciseTypes) == 0 {
		exerciseTypes = make([]model.ExerciseType, 0)
	}

	return cc.JSON(http.StatusOK, exerciseTypes)
}

func Post(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[exerciseTypePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newModel := model.BuildExerciseType(params.UserId, params.Name, params.Note, nil, nil)
	err = cc.ExerciseTypeCrud.Insert(newModel)
	if err != nil {
		return err
	}

	// TODO: Check if this is needed to return
	return cc.JSON(http.StatusOK, newModel)
}

func Put(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[exerciseTypePutPrams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.ExerciseTypeCrud.Update(&model)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
