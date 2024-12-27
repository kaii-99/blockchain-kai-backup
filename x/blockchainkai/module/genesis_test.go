package blockchainkai_test

import (
	"testing"

	keepertest "github.com/kaii-99/blockchain-kai/testutil/keeper"
	"github.com/kaii-99/blockchain-kai/testutil/nullify"
	blockchainkai "github.com/kaii-99/blockchain-kai/x/blockchainkai/module"
	"github.com/kaii-99/blockchain-kai/x/blockchainkai/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlockchainkaiKeeper(t)
	blockchainkai.InitGenesis(ctx, k, genesisState)
	got := blockchainkai.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
