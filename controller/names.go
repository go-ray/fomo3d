package controller

import (
	"net/http"

	"github.com/go-ray/fomo3d/gateway"
	"github.com/go-ray/logging"
	"github.com/gorilla/mux"
)

func NamesHandler(w http.ResponseWriter, r *http.Request) {
	names, err := gateway.GetNames()
	if err != nil {
		logging.Error("get players err:", err)
	}

	w.Write(names)
}

func KeyHolderStatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addr := vars["address"]
	stats, err := gateway.GetKeyHolderStats(addr)
	if err != nil {
		logging.Error("get players err:", err)
	}

	w.Write(stats)
}
