package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"astianetwork/x/token/keeper"
	"astianetwork/x/token/types"
)

func createNCoin(keeper keeper.Keeper, ctx context.Context, n int) []types.Coin {
	items := make([]types.Coin, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].Name = strconv.Itoa(i)
		items[i].Amount = strconv.Itoa(i)
		_ = keeper.Coin.Set(ctx, iu, items[i])
		_ = keeper.CoinSeq.Set(ctx, iu)
	}
	return items
}

func TestCoinQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNCoin(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetCoinRequest
		response *types.QueryGetCoinResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCoinRequest{Id: msgs[0].Id},
			response: &types.QueryGetCoinResponse{Coin: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCoinRequest{Id: msgs[1].Id},
			response: &types.QueryGetCoinResponse{Coin: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCoinRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetCoin(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestCoinQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNCoin(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCoinRequest {
		return &types.QueryAllCoinRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListCoin(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coin), step)
			require.Subset(t, msgs, resp.Coin)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListCoin(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coin), step)
			require.Subset(t, msgs, resp.Coin)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListCoin(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Coin)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListCoin(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
