package gateway

import (
	"time"

	"github.com/go-ray/logging"
)

type ccdata struct {
	t    time.Time
	data []byte
}

var cc = map[string]*ccdata{}
var cstats = map[string][]byte{}

var tmer = time.NewTimer(60 * time.Second)

func update() {
	players, err := getPlayers()
	if err != nil {
		logging.Error("getPlayers from remote failed:", err)
	}
	ccp := &ccdata{
		t:    time.Now(),
		data: players,
	}
	cc[KeyPlayers] = ccp

	names, err := getNames()
	if err != nil {
		logging.Error("getNames from remote failed:", err)
	}
	ccn := &ccdata{
		t:    time.Now(),
		data: names,
	}
	cc[KeyNames] = ccn
}

func Init() {
	cc = make(map[string]*ccdata)
	update()
	go func() {
		for {
			<-tmer.C
			update()
		}
	}()
}

var (
	KeyPlayers = "players"
	KeyNames   = "names"
)

func Players() []byte {
	return cc[KeyPlayers].data
}

func Names() []byte {
	return cc[KeyNames].data
}
