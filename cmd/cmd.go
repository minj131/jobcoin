package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/minj131/jobcoin/config"
	"github.com/minj131/jobcoin/pkg/jobcoin"
	"github.com/minj131/jobcoin/pkg/util"

	"github.com/manifoldco/promptui"
)

var defaultConfig = filepath.Join("config.yaml")

func Jobcoin() error {
	fmt.Println("Welcome to the Jobcoin mixer!")

	for {
		addresses, mixer, err := prompt()
		if err != nil {
			return err
		}

		if addresses == nil && mixer == "" {
			fmt.Println("Goodbye")
			return nil
		}

		cfg := config.New(addresses, mixer)

		err = cfg.ReadConfigFile(defaultConfig)
		if err != nil {
			return err
		}

		err = cfg.Init("debug")
		if err != nil {
			return err
		}

		jobcoin := jobcoin.New(cfg)

		err = jobcoin.Run()
		if err != nil {
			return err
		}
	}
}

func prompt() ([]string, string, error) {
	var addresses []string
	var address string

	prompt := promptui.Prompt{
		Label: "Please enter a comma-separated list of new, unused Jobcoin addresses where your mixed Jobcoins will be sent [blank to quit]:",
		Validate: func(input string) error {
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return nil, "", err
	}

	if result == "" {
		return nil, "", nil
	}

	addresses = strings.Split(strings.TrimSpace(result), ",")
	address = util.GenerateRandomHash(32)
	fmt.Printf("You may now send Jobcoins to address %s. They will be mixed and sent to your destination addresses.\n", address)

	return addresses, address, nil
}
