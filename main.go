// +build linux
package main

import (
	"flag"
	"log"
)

func loadConfig() *Config {
	var configFile string

	flag.StringVar(&configFile, "config", "/etc/go-keyholder.yml", "configuration file")
	flag.Parse()

	cfg := NewDefaultConfig()
	if configFile != "" {
		cfg.ReadFile(configFile)
	}
	return cfg
}

func main() {
	cfg := loadConfig()

	log.Println("listening on", cfg.SocketPath)
	log.Printf("\ttrusted user: uid=%d(%s)", cfg.TrustedUserId, cfg.TrustedUser)
	log.Printf("\tallowed group: gid=%d(%s)", cfg.AllowedGroupId, cfg.AllowedGroup)

	l, err := NewListener(cfg)
	if err != nil {
		log.Fatal(err)
	}

	l.Serve()
}
