package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type API struct {
	AddressUrl      string
	TransactionsUrl string
}

type AddressInfo struct {
	Balance      string          `json:"balance"`
	Transactions []*Transactions `json:"transactions"`
}

type Transactions struct {
	Timestamp   string `json:"timestamp"`
	FromAddress string `json:"fromAddress,omitempty"`
	ToAddress   string `json:"toAddress"`
	Amount      string `json:"amount"`
}

type PostTransactions struct {
	From   string `json:"fromAddress"`
	To     string `json:"toAddress"`
	Amount string `json:"amount"`
}

func New(address, transaction string) *API {
	return &API{
		AddressUrl:      address,
		TransactionsUrl: transaction,
	}
}

func (a *API) GetAddressInfo(address string) (*AddressInfo, error) {
	resp, err := http.Get(a.AddressUrl + address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	addressInfo := &AddressInfo{}
	err = json.Unmarshal(body, addressInfo)
	if err != nil {
		return nil, err
	}

	return addressInfo, nil
}

func (a *API) GetTransactions() ([]Transactions, error) {
	resp, err := http.Get(a.TransactionsUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	transactions := make([]Transactions, 0)
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (a *API) PostTransaction(from, to, amount string) error {
	body := &PostTransactions{
		To:     to,
		From:   from,
		Amount: amount,
	}

	bodybytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = http.Post(a.TransactionsUrl, "application/json", bytes.NewBuffer(bodybytes))

	return err
}
