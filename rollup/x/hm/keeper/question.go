package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"hm/x/hm/types"
)

// SetQuestion set a specific question in the store from its index
func (k Keeper) SetQuestion(ctx context.Context, question types.Question) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKeyPrefix))
	b := k.cdc.MustMarshal(&question)
	store.Set(types.QuestionKey(
		question.Index,
	), b)
}

// GetQuestion returns a question from its index
func (k Keeper) GetQuestion(
	ctx context.Context,
	index string,

) (val types.Question, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKeyPrefix))

	b := store.Get(types.QuestionKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveQuestion removes a question from the store
func (k Keeper) RemoveQuestion(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKeyPrefix))
	store.Delete(types.QuestionKey(
		index,
	))
}

// GetAllQuestion returns all question
func (k Keeper) GetAllQuestion(ctx context.Context) (list []types.Question) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Question
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
