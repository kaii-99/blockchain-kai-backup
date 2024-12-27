package resource

import (
    "github.com/cosmos/cosmos-sdk/types"
    "github.com/kaii-99/blockchain-kai/x/resource/keeper"
    "github.com/kaii-99/blockchain-kai/x/resource/types"
)

func NewQueryHandler(k keeper.Keeper) types.QueryServer {
    return &queryServer{k}
}

type queryServer struct {
    keeper keeper.Keeper
}

func (q *queryServer) QueryAllResource(c context.Context, req *types.QueryAllResourceRequest) (*types.QueryAllResourceResponse, error) {
    return q.keeper.QueryAllResources(types.UnwrapSDKContext(c), req)
}

