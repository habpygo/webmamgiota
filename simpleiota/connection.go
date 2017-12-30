package simpleiota

import "github.com/iotaledger/giota"

func NewConnection(provider, seed string) (*Connection, error) {
	return &Connection{
		api:      giota.NewAPI(provider, nil),
		seed:     seed,
		security: 3,
		mwm:      15,
	}, nil
}

type Connection struct {
	api      *giota.API
	seed     string
	security int
	mwm      int64
}

func (c *Connection) SendToApi(trs []giota.Transfer) (giota.Bundle, error) {
	seed, err := giota.ToTrytes(c.seed)
	if err != nil {
		return nil, err
	}
	_, bestPow := giota.GetBestPoW()
	return giota.Send(c.api, seed, c.security, trs, c.mwm, bestPow)
}

func (c *Connection) FindTransactions(req giota.FindTransactionsRequest) ([]giota.Transaction, error) {
	found, err := c.api.FindTransactions(&req)
	if err != nil {
		return nil, err
	}
	txs, err := c.api.GetTrytes(found.Hashes)
	return txs.Trytes, err
}
