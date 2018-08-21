package config

import ini "gopkg.in/ini.v1"

var (
	// Addr is the listening address
	Addr = ":6060"

	// WorkingDirectory contains where are all the themes, plugins, etc. stored ?
	WorkingDirectory = "."
)

// Load loads the configuration file
func Load() {
	cfg, err := ini.Load("config.ini")

	if err != nil {
		panic(err)
	}

	var sec *ini.Section

	// [http]
	sec = cfg.Section("http")

	Addr = sec.Key("addr").MustString(Addr)

	// [data]
	sec = cfg.Section("data")

	WorkingDirectory = sec.Key("wd").MustString(WorkingDirectory)
}
