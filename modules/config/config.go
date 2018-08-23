package config

import (
	"path/filepath"

	ini "gopkg.in/ini.v1"
)

var (
	// Addr is the listening address
	Addr = ":6060"

	// Db is the database configuration
	Db struct {
		User     string
		Password string
		Name     string
		Addr     string
	}

	// WorkingDirectory contains where are all the themes, plugins, etc. stored ?
	WorkingDirectory = "."

	// Debug ...
	Debug = false
)

// Load loads the configuration file
func Load() {
	cfg, err := ini.Load("data/config.ini")

	if err != nil {
		panic(err)
	}

	Addr = cfg.Section("http").Key("addr").MustString(Addr)

	Db.User = cfg.Section("db").Key("user").MustString("root")
	Db.Password = cfg.Section("db").Key("password").MustString("")
	Db.Name = cfg.Section("db").Key("name").MustString("golog")
	Db.Addr = cfg.Section("db").Key("addr").MustString("root")

	WorkingDirectory, _ = filepath.Abs("data")

	Debug, err = cfg.Section("").Key("debug").Bool()
	if err != nil {
		Debug = false
	}
}
