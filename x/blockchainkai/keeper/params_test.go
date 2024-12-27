package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kaii-99/blockchain-kai/testutil/keeper"
	"github.com/kaii-99/blockchain-kai/x/blockchainkai/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlockchainkaiKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
