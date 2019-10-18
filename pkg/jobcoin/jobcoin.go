package jobcoin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/minj131/jobcoin/config"
	"github.com/minj131/jobcoin/pkg/api"
	"github.com/minj131/jobcoin/pkg/util"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/api.go -fake-name API . API
type API interface {
	GetAddressInfo(address string) (*api.AddressInfo, error)
	GetTransactions() ([]api.Transactions, error)
	PostTransaction(from, to, amount string) error
}

type Jobcoin struct {
	Config *config.Config
	Logger *zap.SugaredLogger
	API    API
}

func New(config *config.Config) *Jobcoin {
	return &Jobcoin{
		Config: config,
		Logger: config.Logger.Sugar().Named("Jobcoin"),
		API:    api.New(config.API.AddressUrl, config.API.TransactionsUrl),
	}
}

func (j *Jobcoin) Run() error {
	sBalance, err := j.PollAddressForChange()
	if err != nil {
		return errors.Wrap(err, "failed to poll for change in mixer address")
	}

	// Transfer to bighouse account
	err = j.TransferJobcoin(j.Config.MixerAddress, j.Config.AdminAddress, sBalance)
	if err != nil {
		return errors.Wrap(err, "failed to transfer coins to admin account")
	}

	fBalance, _ := strconv.ParseFloat(sBalance, 64)
	newBalance, fee := util.FeeCalculator(fBalance, *j.Config.Fee)

	fmt.Printf("Received coins. Fee: %f Jobcoins, the new balance to mix is %f Jobcoins\n", fee, newBalance)
	dividedAmount := util.DivideEvenly(newBalance, len(j.Config.DestAddresses))
	for _, address := range j.Config.DestAddresses {
		fmt.Printf("Sending %f Jobcoins to Address: %s\n", dividedAmount, address)
		err = j.TransferJobcoin(j.Config.AdminAddress, address, fmt.Sprintf("%f", dividedAmount))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to send jobcoins to %s", address))
		}
	}

	return nil
}

// Poll wallet change every 10 seconds
func (j *Jobcoin) PollAddressForChange() (string, error) {
	for {
		fmt.Println("Waiting for coins...")
		time.Sleep(5 * time.Second)
		mixerInfo, err := j.API.GetAddressInfo(j.Config.MixerAddress)
		if err != nil {
			return "", err
		}
		balance := mixerInfo.Balance
		if balance != "0" {
			return balance, nil
		}
	}
}

func (j *Jobcoin) TransferJobcoin(from, to, balance string) error {
	return j.API.PostTransaction(from, to, balance)
}

// API Wrappers
