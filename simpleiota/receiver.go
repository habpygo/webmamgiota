package simpleiota

import (
	"sort"
	"time"

	"github.com/iotaledger/mamgoiota/mamutils" // in $GOPATH/src

	"github.com/iotaledger/giota"
)

type ApiFinder interface {
	FindTransactions(giota.FindTransactionsRequest) ([]giota.Transaction, error)
}

type Transaction struct {
	Message   string
	Value     int64
	Timestamp time.Time
	Recipient string
}

func ReadTransactions(address string, f ApiFinder) ([]Transaction, error) {
	iotaAdress, err := giota.ToAddress(address)
	if err != nil {
		return nil, err
	}

	req := giota.FindTransactionsRequest{
		Addresses: []giota.Address{iotaAdress},
	}

	foundTx, err := f.FindTransactions(req)
	if err != nil {
		return nil, err
	}

	sort.Slice(foundTx, func(i, j int) bool {
		return !(foundTx[i].Timestamp.Unix() < foundTx[j].Timestamp.Unix())
	})

	transactions := make([]Transaction, len(foundTx))
	for i, t := range foundTx {
		//TODO create Trim Util
		//Have to cut off one 9-padding at the end so it can be parsed
		message, err := mamutils.FromMAMTrytes(t.SignatureMessageFragment[:len(t.SignatureMessageFragment)-1])
		if err != nil {
			return nil, err
		}
		transactions[i] = Transaction{
			Message:   message,
			Value:     t.Value,
			Timestamp: t.Timestamp,
			Recipient: string(t.Address),
		}
	}

	return transactions, nil
}
