package controller

import (
	"net/http"

	"github.com/go-ray/fomo3d/gateway"
	"github.com/go-ray/logging"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	players, err := gateway.GetPlayers()
	if err != nil {
		logging.Error("get players err:", err)
	}

	w.Write(players)
}

func Root(w http.ResponseWriter, r *http.Request) {
	logging.Debug("path:", r.URL.Path)
	w.Write([]byte("succes"))
}
