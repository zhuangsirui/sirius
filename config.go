package main

import "github.com/BurntSushi/toml"

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type WebDav struct {
	Prefix  string `toml:"prefix"`
	RootDir string `toml:"root_dir"`
	Users   []User `toml:"users"`
}

type Http struct {
	IP   string `toml:"ip"`
	Port int    `toml:"port"`
}

type Config struct {
	WebDav WebDav `toml:"webdav"`
	Http   Http   `toml:"http"`
}

func (c *Config) ParseFile(file string) error {
	_, err := toml.DecodeFile(file, c)
	return err
}
