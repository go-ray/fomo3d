package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/go-ray/fomo3d/conf"
	"github.com/go-ray/logging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	version   string
	commit    string
	branch    string
	compileAt string
	config    string
)

func main() {
	app := cli.NewApp()
	app.Action = fomoProxy
	app.Name = "fomo-proxy"
	timestamp, _ := strconv.ParseInt(compileAt, 10, 64)
	app.Compiled = time.Unix(timestamp, 0)
	app.Version = fmt.Sprintf("%s\n branch: %s\n commit: %s\n compileAt:%s", version, branch, commit, app.Compiled)

	app.Usage = "the fomo-proxy command line interface"
	app.Copyright = "Copyright 2017-2018 The Authors"

	app.Flags = append(app.Flags, ConfigFlag)

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Commands = []cli.Command{}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func fomoProxy(ctx *cli.Context) error {
	start(ctx)
	return nil
}

func start(ctx *cli.Context) {
	cfg, err := loadConfig(ctx)
	if err != nil {
		panic(err)
	}

	conf.Cfg = cfg
	setLog(conf.Cfg)

	setPprof(conf.Cfg)
	setDatabase(conf.Cfg)

	wait := ctx.Duration(WaitFlag.Name)

	r := mux.NewRouter()
	setHandles(r)
	r.Use(loggingMiddleware)

	srv := &http.Server{
		Addr:         conf.Cfg.Server.Addr,
		WriteTimeout: conf.Cfg.Server.WriteTimeout,
		ReadTimeout:  conf.Cfg.Server.ReadTimeout,
		IdleTimeout:  conf.Cfg.Server.IdleTimeout,
		Handler:      handlers.CORS()(r),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logging.Println(err)
		}
	}()
	logging.Info("server start and listening on:", conf.Cfg.Server.Addr)
	logging.Info("server writetimeout:", conf.Cfg.Server.WriteTimeout)
	logging.Info("server readtimeout:", conf.Cfg.Server.ReadTimeout)
	logging.Info("server idletimeout:", conf.Cfg.Server.IdleTimeout)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Create a deadline to wait for.
	contxt, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(contxt)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logging.Println("shutting down")
	os.Exit(0)

}
