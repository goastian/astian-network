package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v10/modules/core/24-host"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		PortId: PortID, CoinList: []Coin{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	coinIdMap := make(map[uint64]bool)
	coinCount := gs.GetCoinCount()
	for _, elem := range gs.CoinList {
		if _, ok := coinIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for coin")
		}
		if elem.Id >= coinCount {
			return fmt.Errorf("coin id should be lower or equal than the last id")
		}
		coinIdMap[elem.Id] = true
	}

	return gs.Params.Validate()
}
