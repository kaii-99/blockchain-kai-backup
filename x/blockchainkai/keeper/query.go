package keeper

import (
	"github.com/kaii-99/blockchain-kai/x/blockchainkai/types"
)

var _ types.QueryServer = Keeper{}
