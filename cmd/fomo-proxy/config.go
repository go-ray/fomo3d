package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"

	"github.com/go-ray/fomo3d/conf"
	"github.com/go-ray/fomo3d/database"
	"github.com/go-ray/logging"
	cli "gopkg.in/urfave/cli.v1"
)

func loadConfig(ctx *cli.Context) (*conf.Config, error) {
	cfg := &conf.Config{}
	configpath := config
	data, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

func setLog(cfg *conf.Config) {
	logging.InitLogger(cfg.Log.Path, cfg.Log.Name, cfg.Log.Level, cfg.Log.Format)
}

func setPprof(cfg *conf.Config) {
	if cfg.Pprof.Enable {
		go func() {
			logging.Debug("running http pprof on:", cfg.Pprof.Host)
			http.HandleFunc("/debug/version", getVersion)
			http.ListenAndServe(cfg.Pprof.Host, nil)
		}()
	}
}

func setDatabase(cfg *conf.Config) {
	for k, c := range cfg.DBs {
		if !c.Enable {
			continue
		}

		err := database.InitDatabaseConfig(k, c)
		if err != nil {
			logging.Error("Open ["+k+"] Database err:", err)
			panic(err)
		}
	}
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	i, _ := strconv.ParseInt(compileAt, 10, 64)
	t := time.Unix(i, 0)

	v := &conf.Version{
		Version:   version,
		Commit:    commit,
		Branch:    branch,
		CompileAt: t.String(),
	}
	d, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(d)
}
