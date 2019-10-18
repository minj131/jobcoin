package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	API           *API `yaml:"api"`
	DestAddresses []string
	MixerAddress  string
	Fee           *float64 `yaml:"fee"`
	AdminAddress  string   `yaml:"adminAccount"`
}

type API struct {
	BaseUrl         string `yaml:"baseUrl"`
	AddressUrl      string `yaml:"addressUrl"`
	TransactionsUrl string `yaml:"transactionsUrl"`
}

func New(addresses []string, mixer string) *Config {
	return &Config{
		DestAddresses: addresses,
		MixerAddress:  mixer,
	}
}

func (c *Config) Init(level string) error {
	if c.API == nil {
		return errors.New("API config is not passed")
	}

	if c.DestAddresses == nil || len(c.DestAddresses) == 0 {
		return errors.New("Destination addresses must be set and not empty")
	}

	if c.Fee == nil {
		f := float64(0.05)
		c.Fee = &f
	}

	if c.AdminAddress == "" {
		return errors.New("Admin address is not passed")
	}

	if level == "" {
		level = "debug"
	}

	return nil
}

func (c *Config) ReadConfigFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "failed to get api configurations")
	}

	err = yaml.Unmarshal([]byte(file), c)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal api configurations")
	}

	return nil
}
