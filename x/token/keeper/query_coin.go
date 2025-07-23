package keeper

import (
	"context"
	"errors"

	"astianetwork/x/token/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListCoin(ctx context.Context, req *types.QueryAllCoinRequest) (*types.QueryAllCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coins, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Coin,
		req.Pagination,
		func(_ uint64, value types.Coin) (types.Coin, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinResponse{Coin: coins, Pagination: pageRes}, nil
}

func (q queryServer) GetCoin(ctx context.Context, req *types.QueryGetCoinRequest) (*types.QueryGetCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coin, err := q.k.Coin.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetCoinResponse{Coin: coin}, nil
}
