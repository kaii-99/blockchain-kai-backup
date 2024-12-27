package keeper

import (
	"fmt"

	"cosmossdk.io/log"
        "cosmossdk.io/core/store"
        "cosmossdk.io/store/prefix"
        "github.com/cosmos/cosmos-sdk/types"
        "github.com/cosmos/cosmos-sdk/codec"
	"github.com/kaii-99/blockchain-kai/x/resource/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		storeKey     sdk.StoreKey

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	storeKey sdk.StoreKey,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		storeKey:     storeKey,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// QueryAllResources handles the query for fetching all resources with optional filters
func (k Keeper) QueryAllResources(ctx sdk.Context, req *types.QueryAllResourceRequest) (*types.QueryAllResourceResponse, error) {
    store := ctx.KVStore(k.storeKey)
    resourceStore := prefix.NewStore(store, types.KeyPrefixResource)

    var resources []types.Resource
    iterator := resourceStore.Iterator(nil, nil)

    // Iterate through the stored resources and filter based on parameters
    for ; iterator.Valid(); iterator.Next() {
        var resource types.Resource
        k.cdc.MustUnmarshal(iterator.Value(), &resource)

        // Apply filtering if required
        if req.Owner != "" && resource.Owner == req.Owner {
            resources = append(resources, resource)
        } else if req.Name != "" && resource.Name == req.Name {
            resources = append(resources, resource)
        } else if req.Owner == "" && req.Name == "" {
            resources = append(resources, resource)
        }
    }
    iterator.Close()

    return &types.QueryAllResourceResponse{Resources: resources}, nil
}
