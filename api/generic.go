package api

import (
	"net/http"
	"strconv"
	"trainer-helper/store"
)

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
