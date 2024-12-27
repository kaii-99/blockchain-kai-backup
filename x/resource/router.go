package resource

import (
    "github.com/cosmos/cosmos-sdk/types"
    "github.com/kaii-99/blockchain-kai/x/resource/keeper"
    "github.com/kaii-99/blockchain-kai/x/resource/types"
)

func NewQueryHandler(k keeper.Keeper) types.QueryServer {
    return &queryServer{k}
}

// Register the query with the router
func RegisterQueryHandler(cdc *codec.Codec, queryRouter types.QueryRouter, keeper keeper.Keeper) {
    queryRouter.AddRoute(types.QueryAllResource, NewQueryHandler(keeper))
}

