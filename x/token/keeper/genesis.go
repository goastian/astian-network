package keeper

import (
	"context"
	"errors"

	"astianetwork/x/token/types"

	"cosmossdk.io/collections"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	if err := k.Port.Set(ctx, genState.PortId); err != nil {
		return err
	}
	for _, elem := range genState.CoinList {
		if err := k.Coin.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.CoinSeq.Set(ctx, genState.CoinCount); err != nil {
		return err
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.PortId, err = k.Port.Get(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}
	err = k.Coin.Walk(ctx, nil, func(key uint64, elem types.Coin) (bool, error) {
		genesis.CoinList = append(genesis.CoinList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.CoinCount, err = k.CoinSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
