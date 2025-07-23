package keeper

import (
	"context"
	"errors"
	"fmt"

	"astianetwork/x/token/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCoin(ctx context.Context, msg *types.MsgCreateCoin) (*types.MsgCreateCoinResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	nextId, err := k.CoinSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var coin = types.Coin{
		Id:      nextId,
		Creator: msg.Creator,
		Name:    msg.Name,
		Amount:  msg.Amount,
	}

	if err = k.Coin.Set(
		ctx,
		nextId,
		coin,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coin")
	}

	return &types.MsgCreateCoinResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdateCoin(ctx context.Context, msg *types.MsgUpdateCoin) (*types.MsgUpdateCoinResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var coin = types.Coin{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Amount:  msg.Amount,
	}

	// Checks that the element exists
	val, err := k.Coin.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coin")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Coin.Set(ctx, msg.Id, coin); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update coin")
	}

	return &types.MsgUpdateCoinResponse{}, nil
}

func (k msgServer) DeleteCoin(ctx context.Context, msg *types.MsgDeleteCoin) (*types.MsgDeleteCoinResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Coin.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coin")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Coin.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete coin")
	}

	return &types.MsgDeleteCoinResponse{}, nil
}
