package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "game"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_game"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	PrefixWhitelistedContract = []byte{0x01}
	PrefixAccountDeposit      = []byte{0x02}
)

func AccountDepositKey(addr sdk.AccAddress) []byte {
	return append(PrefixAccountDeposit, addr...)
}
