package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Fqdn    string
	Address string
	Port    int
	TTL     int
	NS      string
	NSIp    string
)

func LoadConfig(debug bool) {
	// load config file
	cfg, err := ini.Load("/etc/wildip-resolver-dns/config.ini")
	if debug {
		// load config file from current directory when debugging
		fmt.Println("DEBUG: Loading config from ./config.ini")
		cfg, err = ini.Load("config.ini")
	}
	if err != nil {
		panic(err)
	}

	// read values
	Fqdn = cfg.Section("server").Key("fqdn").String()
	Address = cfg.Section("server").Key("address").String()
	Port = cfg.Section("server").Key("port").MustInt(53)
	TTL = cfg.Section("dns").Key("ttl").MustInt(600)
	NS = cfg.Section("dns").Key("ns").String()
	NSIp = cfg.Section("dns").Key("nsIp").String()
}
