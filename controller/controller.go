package controller

import (
	"net/http"
	"strconv"

	"github.com/go-ray/logging"
	"github.com/gorilla/mux"
)

func getoffamount(r *http.Request) (ioff, ia int) {
	vars := mux.Vars(r)
	offset := vars["offset"]
	amount := vars["amount"]

	ioff, err := strconv.Atoi(offset)
	if err != nil {
		logging.Error("players offset incorrect:", err)
	}
	ia, err = strconv.Atoi(amount)
	if err != nil {
		logging.Error("players amount incorrect:", err)
	}
	return
}
