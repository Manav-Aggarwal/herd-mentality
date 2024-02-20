package types

const (
	// ModuleName defines the module name
	ModuleName = "hm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hm"
)

var (
	ParamsKey = []byte("p_hm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
