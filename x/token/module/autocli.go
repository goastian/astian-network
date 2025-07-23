package token

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"astianetwork/x/token/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
			RpcMethod: "ListCoin",
			Use: "list-coin",
			Short: "List all coin",
		},
		{
			RpcMethod: "GetCoin",
			Use: "get-coin [id]",
			Short: "Gets a coin by id",
			Alias: []string{"show-coin"},
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
			RpcMethod: "CreateCoin",
			Use: "create-coin [name]",
			Short: "Create coin",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
		},
		{
			RpcMethod: "UpdateCoin",
			Use: "update-coin [id] [name]",
			Short: "Update coin",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}},
		},
		{
			RpcMethod: "DeleteCoin",
			Use: "delete-coin [id]",
			Short: "Delete coin",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
