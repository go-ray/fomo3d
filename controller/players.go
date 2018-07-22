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
	Data  []pd `json:"data,omitempty"`
	Total int  `json:"total,omitempty"`
}

type pd struct {
	A   string `json:"a,omitempty"`
	I   string `json:"i,omitempty"`
	K   string `json:"k,omitempty"`
	P   int    `json:"p,omitempty"`
	R   string `json:"r,omitempty"`
	Re  string `json:"re,omitempty"`
	Ti  int    `json:"ti,omitempty"`
	Tr  int    `json:"tr,omitempty"`
	Tre int    `json:"tre,omitempty"`
}
