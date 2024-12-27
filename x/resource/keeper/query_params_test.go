package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kaii-99/blockchain-kai/testutil/keeper"
	"github.com/kaii-99/blockchain-kai/x/resource/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.ResourceKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
