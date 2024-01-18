package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// go install github.com/BurntSushi/toml/cmd/tomlv@latest

func LoadFromFile(filepath string) error {
	c := DefaultConfig()
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

func LoadFromEnv() error {
	c := DefaultConfig()
	if err := env.Parse(c); err != nil {
		return err
	}
	config = c
	return nil
}
