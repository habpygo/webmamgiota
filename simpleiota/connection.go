package simpleiota

import "github.com/iotaledger/giota"

func NewConnection(provider, seed string) (*Connection, error) {
	seedTrytes, err := giota.ToTrytes(seed)
	if err != nil {
		return nil, err
	}
	return &Connection{
		api:      giota.NewAPI(provider, nil),
		seed:     seedTrytes,
		security: 3,
		mwm:      15,
	}, nil
}

type Connection struct {
	api      *giota.API
	seed     giota.Trytes
	security int
	mwm      int64
}

func (c *Connection) SendToApi(trs []giota.Transfer) (giota.Bundle, error) {
	_, bestPow := giota.GetBestPoW()
	return giota.Send(c.api, c.seed, c.security, trs, c.mwm, bestPow)
}
