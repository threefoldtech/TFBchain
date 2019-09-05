package main

import (
	"github.com/threefoldtech/rivine/types"
	tfbchaintypes "github.com/threefoldtech/TFBchain/pkg/types"
	"github.com/threefoldtech/rivine/extensions/minting"
	mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"
	

	"github.com/threefoldtech/rivine/pkg/client"
)

func RegisterDevnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

func RegisterTestnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}


func registerTransactions(cli *client.CommandLineClient) {
	// create minting plugin client...
	mintingCLI := mintingcli.NewPluginConsensusClient(cli)
	// ...and register minting types
	types.RegisterTransactionVersion(tfbchaintypes.TransactionVersionMinterDefinition, minting.MinterDefinitionTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  tfbchaintypes.TransactionVersionMinterDefinition,
	})
	types.RegisterTransactionVersion(tfbchaintypes.TransactionVersionCoinCreation, minting.CoinCreationTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  tfbchaintypes.TransactionVersionCoinCreation,
	})
	types.RegisterTransactionVersion(tfbchaintypes.TransactionVersionCoinDestruction, minting.CoinDestructionTransactionController{
		TransactionVersion: tfbchaintypes.TransactionVersionCoinDestruction,
	})

	
}
