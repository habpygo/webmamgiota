package simpleiota

import (
	"testing"
	"time"

	"github.com/iotaledger/mamgoiota/mamutils"

	"github.com/iotaledger/giota"
	"github.com/stretchr/testify/assert"
)

type FakeFinder struct {
	Find func(giota.FindTransactionsRequest) ([]giota.Transaction, error)
}

func (f *FakeFinder) FindTransactions(req giota.FindTransactionsRequest) ([]giota.Transaction, error) {
	return f.Find(req)
}

func TestReadTransactions(t *testing.T) {
	assert := assert.New(t)

	account, err := giota.NewAddress(giota.NewSeed(), 0, 1)
	assert.Nil(err)

	txs, err := ReadTransactions(string(account), &FakeFinder{
		Find: func(req giota.FindTransactionsRequest) ([]giota.Transaction, error) {
			assert.Len(req.Addresses, 1)
			assert.EqualValues(account, req.Addresses[0])

			message, err := mamutils.ToMAMTrytes("Test")
			assert.Nil(err)
			return []giota.Transaction{
				giota.Transaction{
					SignatureMessageFragment: message + "9",
					Value:     1000,
					Timestamp: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
					Address:   "Recipient",
				},
			}, nil
		},
	})

	assert.Nil(err)
	assert.Len(txs, 1)
	assert.EqualValues(Transaction{
		Message:   "Test",
		Value:     1000,
		Timestamp: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
		Recipient: "Recipient",
	}, txs[0])
}
