package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"hm/x/hm/types"
)

// SetAnswer set a specific answer in the store from its index
func (k Keeper) SetAnswer(ctx context.Context, answer types.Answer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AnswerKeyPrefix))
	b := k.cdc.MustMarshal(&answer)
	store.Set(types.AnswerKey(
		answer.Index,
	), b)
}

// GetAnswer returns a answer from its index
func (k Keeper) GetAnswer(
	ctx context.Context,
	index string,

) (val types.Answer, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AnswerKeyPrefix))

	b := store.Get(types.AnswerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAnswer removes a answer from the store
func (k Keeper) RemoveAnswer(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AnswerKeyPrefix))
	store.Delete(types.AnswerKey(
		index,
	))
}

// GetAllAnswer returns all answer
func (k Keeper) GetAllAnswer(ctx context.Context) (list []types.Answer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AnswerKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Answer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
