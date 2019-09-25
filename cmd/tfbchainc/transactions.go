package main

import (
	tfbchaintypes "github.com/threefoldtech/TFBchain/pkg/types"
	"github.com/threefoldtech/rivine/extensions/minting"
	mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"
	"github.com/threefoldtech/rivine/types"

	"github.com/threefoldtech/rivine/pkg/client"
)

func RegisterDevnetTransactions(bc *client.BaseClient) {
	registerTransactions(bc)
}

func RegisterTestnetTransactions(bc *client.BaseClient) {
	registerTransactions(bc)
}

func registerTransactions(bc *client.BaseClient) {
	// create minting plugin client...
	mintingCLI := mintingcli.NewPluginConsensusClient(bc)
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
