package exercise_type

import (
	"errors"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)
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
	cc := c.(*api.DbContext)

	params, err := api.BindParams[exerciseTypePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newModel := model.BuildExerciseType(cc.Claims.Subject, params.Name, params.Note, params.YoutubeLink, params.FilePath, params.MediaType)
	err = cc.ExerciseTypeCrud.Insert(newModel)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newModel)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PutModel[exerciseTypePutPrams](cc, cc.ExerciseTypeCrud)
}
