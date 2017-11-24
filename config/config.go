package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

const (
	DefaultDutyCount = 1
)

type Config struct {
	DutyCount      int
	Members        []string
	CustomHolidays []string
}

func Load(configPath string) (Config, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		return Config{}, err
	}

	if config.DutyCount <= 0 {
		config.DutyCount = DefaultDutyCount
	}

	return config, nil
}
