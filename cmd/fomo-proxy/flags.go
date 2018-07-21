package main

import (
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	// ConfigFlag config file path
	ConfigFlag = cli.StringFlag{
		Name:        "config, c",
		Usage:       "load configuration from `FILE`",
		Value:       "config.json",
		Destination: &config,
	}
	WaitFlag = cli.DurationFlag{
		Name:  "timeout",
		Usage: "the duration wait for existing connections to finish e.g 15s or 1m",
		Value: time.Second * 15,
	}
)
