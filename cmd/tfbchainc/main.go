package main

import (
	"fmt"
	"os"

	"github.com/threefoldtech/rivine/pkg/cli"
	"github.com/threefoldtech/rivine/pkg/daemon"

	"github.com/threefoldtech/TFBchain/pkg/config"

	"github.com/threefoldtech/TFBchain/pkg/types"
	mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"

	"github.com/threefoldtech/rivine/modules"
	"github.com/threefoldtech/rivine/pkg/client"
)

func main() {
	// create cli
	bchainInfo := config.GetBlockchainInfo()
	cliClient, err := NewCommandLineClient("http://localhost:21110", bchainInfo.Name, daemon.RivineUserAgent)
	if err != nil {
		panic(err)
	}

	// register goldchain-specific explorer commands
	mintingcli.CreateExploreCmd(cliClient.CommandLineClient)
	mintingcli.CreateConsensusCmd(cliClient.CommandLineClient)

	// add cli wallet extension commands
	mintingcli.CreateWalletCmds(
		cliClient.CommandLineClient,
		types.TransactionVersionMinterDefinition,
		types.TransactionVersionCoinCreation,
		&mintingcli.WalletCmdsOpts{
			CoinDestructionTxVersion: types.TransactionVersionCoinDestruction,
		},
	)

	// define preRun function
	cliClient.PreRunE = func(cfg *client.Config) (*client.Config, error) {
		if cfg == nil {
			bchainInfo := config.GetBlockchainInfo()
			chainConstants := config.GetDefaultGenesis()
			daemonConstants := modules.NewDaemonConstants(bchainInfo, chainConstants)
			newCfg := client.ConfigFromDaemonConstants(daemonConstants)
			cfg = &newCfg
		}

		switch cfg.NetworkName {

		case config.NetworkNameDevnet:
			RegisterDevnetTransactions(cliClient.CommandLineClient)
			cfg.GenesisBlockTimestamp = 1566295200 // timestamp of block #1

		case config.NetworkNameTestnet:
			RegisterTestnetTransactions(cliClient.CommandLineClient)
			cfg.GenesisBlockTimestamp = 1566295200 // timestamp of block #1

		default:
			return nil, fmt.Errorf("Network name %q not recognized", cfg.NetworkName)
		}

		return cfg, nil
	}

	// start cli
	if err := cliClient.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "client exited with an error: ", err)
		// Since no commands return errors (all commands set Command.Run instead of
		// Command.RunE), Command.Execute() should only return an error on an
		// invalid command or flag. Therefore Command.Usage() was called (assuming
		// Command.SilenceUsage is false) and we should exit with exitCodeUsage.
		os.Exit(cli.ExitCodeUsage)
	}
}
