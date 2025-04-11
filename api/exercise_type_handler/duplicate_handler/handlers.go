package exercise_type_duplicate_handler

// import (
// 	"github.com/labstack/echo/v4"
// )

// func Post(c echo.Context) error {
// cc := c.(*schemas.DbContext)
//
// params, err := api.BindParams[exerciseTypeDuplicatePostParams](cc)
// if err != nil {
// 	return cc.BadRequest(err)
// }
//
// newModel := model.BuildExerciseType(params.UserId, params.Name, params.Note, nil, nil)
// err = cc.ExerciseTypeCrud.Insert(newModel)
// if err != nil {
// 	return err
// }
//
// // TODO: Check if this is needed to return
// return cc.JSON(http.StatusOK, newModel)
// }
