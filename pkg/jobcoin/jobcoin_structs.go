package jobcoin

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
