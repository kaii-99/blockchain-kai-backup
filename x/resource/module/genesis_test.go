package resource_test

import (
	"testing"

	keepertest "github.com/kaii-99/blockchain-kai/testutil/keeper"
	"github.com/kaii-99/blockchain-kai/testutil/nullify"
	resource "github.com/kaii-99/blockchain-kai/x/resource/module"
	"github.com/kaii-99/blockchain-kai/x/resource/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourceKeeper(t)
	resource.InitGenesis(ctx, k, genesisState)
	got := resource.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
