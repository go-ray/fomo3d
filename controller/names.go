package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-ray/fomo3d/gateway"
	"github.com/go-ray/logging"
	"github.com/gorilla/mux"
)

func NamesHandler(w http.ResponseWriter, r *http.Request) {
	ioff, ia := getoffamount(r)
	names := gateway.GetNames()
	nrp := Nresp{}
	err := json.Unmarshal(names, &nrp)
	if err != nil {
		logging.Error("unmarshal nresp failed:", err)
	}

	nlen := len(nrp.Data)
	if ioff >= nlen {
		ioff = 0
	}
	ia += ioff
	if ia >= nlen {
		ia = nlen - 1
	}
	if ioff == 0 && ia == 0 {
		ia = nlen - 1
	}
	rnrp := Nresp{
		Data:  nrp.Data[ioff:ia],
		Total: len(nrp.Data),
	}

	rdata, err := json.Marshal(&rnrp)
	if err != nil {
		logging.Error("marshal nresp failed:", err)
	}

	w.Write(rdata)
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

type Nresp struct {
	Data  []nd `json:"data"`
	Total int  `json:"total"`
}

type nd struct {
	Addr  string `json:"addr"`
	Fomol bool   `json:"fomol"`
	Name  string `json:"name"`
	Pid   int    `json:"pid"`
}
