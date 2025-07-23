package types_test

import (
	"testing"

	"astianetwork/x/token/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId:   types.PortID,
				CoinList: []types.Coin{{Id: 0}, {Id: 1}}, CoinCount: 2,
			}, valid: true,
		}, {
			desc: "duplicated coin",
			genState: &types.GenesisState{
				CoinList: []types.Coin{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		}, {
			desc: "invalid coin count",
			genState: &types.GenesisState{
				CoinList: []types.Coin{
					{
						Id: 1,
					},
				},
				CoinCount: 0,
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
