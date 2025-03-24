package exercise_duplicate_handler

// import (
// 	"log"
// 	"trainer-helper/api"
// 	"trainer-helper/model"
//
// 	"github.com/labstack/echo/v4"
// )
//
// func Post(c echo.Context) error {
//
// 	cc := c.(*api.DbContext)
// 	params, err := api.BindParams[exerciseDuplicatePostParams](cc)
// 	if err != nil {
// 		return cc.BadRequest(err)
// 	}
//
// 	copyTimeslot, err := cc.CRUDTimeslot.GetById(params.CopyTimeslotId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// TODO: call get
// 	err = cc.CRUDTimeslot.Update(&model.Timeslot{
// 		IdModel: model.IdModel{
// 			Id: copyTimeslot.Id,
// 		},
// 		Name: copyTimeslot.Name,
// 	})
//
// }
