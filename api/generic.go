package api

import (
	"net/http"
	"strconv"
	"trainer-helper/store"
)

type Modulator[M any] interface {
	ToModel() M
}

func DeleteModel[M any](cc *DbContext, crud store.StoreBase[M]) error {
	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	err = crud.Delete(id)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func PostUndeleteModel[M any](cc *DbContext, crud store.StoreBase[M]) error {
	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	err = crud.UndeleteMany([]int{id})
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func PutModel[R Modulator[M], M any](cc *DbContext, crud store.StoreBase[M]) error {
	params, err := BindParams[R](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.ToModel()
	err = crud.Update(&model)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func PostModel[R Modulator[M], M any](cc *DbContext, crud store.StoreBase[M]) error {
	params, err := BindParams[R](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newModel := params.ToModel()
	err = crud.Insert(&newModel)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newModel)
}
