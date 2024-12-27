package types

const (
	// ModuleName defines the module name
	ModuleName = "blockchainkai"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blockchainkai"
)

var (
	ParamsKey = []byte("p_blockchainkai")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
