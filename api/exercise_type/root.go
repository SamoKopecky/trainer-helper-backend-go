package exercise_type

import (
	"errors"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"trainer-helper/api"
	"trainer-helper/media"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
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

	newModel := model.BuildExerciseType(cc.Claims.Subject, params.Name, params.Note, params.YoutubeLink, params.FilePath, nil, params.MediaType)
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

func PostMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	fileId, err := media.SaveFile(file, cc.Config.MediaFileRepository)
	if err != nil {
		return err
	}

	err = cc.ExerciseTypeCrud.UpdateMediaFile(id, fileId, file.Filename)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func GetMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	model, err := cc.ExerciseTypeCrud.GetById(id)
	if err != nil {
		return cc.BadRequest(err)
	}
	if model.FilePath == nil {
		return cc.NoContent(http.StatusOK)
	}

	fileData, err := os.ReadFile(filepath.Join(cc.Config.MediaFileRepository, *model.FilePath))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read file")
	}

	mediaType := mime.TypeByExtension(filepath.Ext(*model.OriginalFileName))
	return cc.Blob(http.StatusOK, mediaType, fileData)
}

// TODO: Add notifaction when file upload succesful
