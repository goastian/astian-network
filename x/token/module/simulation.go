package token

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"

	tokensimulation "astianetwork/x/token/simulation"
	"astianetwork/x/token/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenGenesis := types.GenesisState{
		Params:   types.DefaultParams(),
		PortId:   types.PortID,
		CoinList: []types.Coin{{Id: 0, Creator: sample.AccAddress()}, {Id: 1, Creator: sample.AccAddress()}}, CoinCount: 2,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateCoin          = "op_weight_msg_token"
		defaultWeightMsgCreateCoin int = 100
	)

	var weightMsgCreateCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCoin, &weightMsgCreateCoin, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCoin = defaultWeightMsgCreateCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCoin,
		tokensimulation.SimulateMsgCreateCoin(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateCoin          = "op_weight_msg_token"
		defaultWeightMsgUpdateCoin int = 100
	)

	var weightMsgUpdateCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCoin, &weightMsgUpdateCoin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoin = defaultWeightMsgUpdateCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCoin,
		tokensimulation.SimulateMsgUpdateCoin(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteCoin          = "op_weight_msg_token"
		defaultWeightMsgDeleteCoin int = 100
	)

	var weightMsgDeleteCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteCoin, &weightMsgDeleteCoin, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCoin = defaultWeightMsgDeleteCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCoin,
		tokensimulation.SimulateMsgDeleteCoin(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
