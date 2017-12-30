package mamutils

import (
	"testing"

	"github.com/iotaledger/giota"
	"github.com/stretchr/testify/assert"
)

func TestToMAMTrytes(t *testing.T) {
	assert := assert.New(t)

	expectedTrytes, err := giota.ToTrytes("RBTC9D9DCDEAFCCDFD9DSC")
	assert.Nil(err)

	tr, err := ToMAMTrytes("Hello World")
	assert.Nil(err)
	assert.EqualValues(expectedTrytes, tr)
}

func TestFromMAMTrytes(t *testing.T) {
	assert := assert.New(t)

	input, err := giota.ToTrytes("RBTC9D9DCDEAFCCDFD9DSC")
	assert.Nil(err)

	m, err := FromMAMTrytes(input)
	assert.Nil(err)
	assert.EqualValues("Hello World", m)
}
