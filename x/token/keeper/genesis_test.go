package keeper_test

import (
	"testing"

	"astianetwork/x/token/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:    types.DefaultParams(),
		PortId:    types.PortID,
		CoinList:  []types.Coin{{Id: 0}, {Id: 1}},
		CoinCount: 2,
	}
	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.CoinList, got.CoinList)
	require.Equal(t, genesisState.CoinCount, got.CoinCount)

}
