package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-ray/fomo3d/gateway"
	"github.com/go-ray/logging"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	ioff, ia := getoffamount(r)

	players := gateway.GetPlayers()
	prp := Presp{}
	err := json.Unmarshal(players, &prp)
	if err != nil {
		logging.Error("unmarshal presp failed:", err)
	}

	plen := len(prp.Data)
	if ioff >= plen {
		ioff = 0
	}
	ia += ioff
	if ia >= plen {
		ia = plen - 1
	}
	if ioff == 0 && ia == 0 {
		ia = plen - 1
	}

	rprp := Presp{
		Data:  prp.Data[ioff:ia],
		Total: len(prp.Data),
	}

	rdata, err := json.Marshal(&rprp)
	if err != nil {
		logging.Error("unmarshal presp failed:", err)
	}
	w.Write(rdata)
}

func Root(w http.ResponseWriter, r *http.Request) {
	logging.Debug("path:", r.URL.Path)
	w.Write([]byte("succes"))
}

type Presp struct {
	Data  []pd `json:"data"`
	Total int  `json:"total"`
}

type pd struct {
	A   string `json:"a"`
	I   string `json:"i"`
	K   string `json:"k"`
	P   int    `json:"p"`
	R   string `json:"r"`
	Re  string `json:"re"`
	Ti  int    `json:"ti"`
	Tr  int    `json:"tr"`
	Tre int    `json:"tre"`
}
