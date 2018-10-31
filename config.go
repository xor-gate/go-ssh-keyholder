package main

import (
	"os"
	"os/user"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TrustedUser    string `yaml:"trusted_user,omitempty"`
	TrustedUserId  uint32
	AllowedGroup   string `yaml:"allowed_group,omitempty"`
	AllowedGroupId uint32
	SocketPath     string `yaml:"socket_path,omitempty"`
}

func NewDefaultConfig() *Config {
	cfg := &Config{
		TrustedUser:  "0",
		AllowedGroup: "0",
		SocketPath:   "/var/run/go-keyholder.sock",
	}
	return cfg
}

func (c *Config) ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	if err = dec.Decode(c); err != nil {
		return err
	}

	if c.TrustedUser != "" {
		if uid, err := strconv.ParseUint(c.TrustedUser, 10, 32); err == nil {
			user, err := user.LookupId(c.TrustedUser)
			if err != nil {
				return err
			}
			c.TrustedUser = user.Name
			c.TrustedUserId = uint32(uid)
		} else {
			user, err := user.Lookup(c.TrustedUser)
			if err != nil {
				return err
			}
			uid, err := strconv.ParseUint(user.Uid, 10, 32)
			if err != nil {
				return err
			}
			c.TrustedUserId = uint32(uid)
		}
	}

	if c.AllowedGroup != "" {
		if gid, err := strconv.ParseUint(c.AllowedGroup, 10, 32); err == nil {
			group, err := user.LookupGroupId(c.AllowedGroup)
			if err != nil {
				return err
			}
			c.AllowedGroup = group.Name
			c.AllowedGroupId = uint32(gid)
		} else {
			group, err := user.LookupGroup(c.AllowedGroup)
			if err != nil {
				return err
			}
			gid, err := strconv.ParseUint(group.Gid, 10, 32)
			if err != nil {
				return err
			}
			c.AllowedGroupId = uint32(gid)
		}
	}

	return nil
}
