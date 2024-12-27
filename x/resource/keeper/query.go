package keeper

import (
	"github.com/kaii-99/blockchain-kai/x/resource/types"
)

var _ types.QueryServer = Keeper{}
