package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

type Config struct {
	API           *API `yaml:"api"`
	DestAddresses []string
	MixerAddress  string
	Fee           *float64 `yaml:"fee"`
	AdminAddress  string   `yaml:"adminAccount"`
	Logger        *zap.Logger
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
	var err error

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

	c.Logger, err = setupLogging(level)
	if err != nil {
		return errors.Wrapf(err, "failed to set up logging")
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

func setupLogging(loglevel string) (*zap.Logger, error) {
	// set up logging
	var level zapcore.Level
	err := level.Set(loglevel)
	if err != nil {
		return nil, err
	}
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zapConfig.Build()
	defer logger.Sync()
	if err != nil {
		return nil, err
	}
	// redirect uses of standard logger
	zap.RedirectStdLog(logger)

	return logger, nil
}
