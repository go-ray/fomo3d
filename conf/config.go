package conf

import "time"

var Cfg *Config

type DBConf struct {
	DriverName   string `json:"drivername,omitempty"`
	DataSource   string `json:"datasource,omitempty"`
	MaxIdleConns int    `json:"maxidleconns,omitempty"`
	MaxOpenConns int    `json:"maxopenconns,omitempty"`
	Enable       bool   `json:"enable,omitempty"`
}

type Config struct {
	FomoApi struct {
		Host string `json:"host,omitempty"`
	}
	Log struct {
		Level  string `json:"level,omitempty"`
		Path   string `json:"path,omitempty"`
		Name   string `json:"file,omitempty"`
		Format string `json:"format,omitempty"`
	}
	Server struct {
		Addr         string        `json:"addr,omitempty"`
		WriteTimeout time.Duration `json:"writeTimeout,omitempty"`
		ReadTimeout  time.Duration `json:"readTimeout,omitempty"`
		IdleTimeout  time.Duration `json:"idleTimeout,omitempty"`
	}
	Pprof struct {
		Enable bool   `json:"enable,omitempty"`
		Host   string `json:"host,omitempty"`
	}
	DBs map[string]DBConf `json:"dbs,omitempty"`
}

type Version struct {
	Version   string `json:"Version,omitempty"`
	Commit    string `json:"Commit,omitempty"`
	Branch    string `json:"Branch,omitempty"`
	CompileAt string `json:"CompileAt,omitempty"`
}
