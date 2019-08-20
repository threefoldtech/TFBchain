package config

import (
	"fmt"
	"math/big"

	"github.com/threefoldtech/rivine/build"
	"github.com/threefoldtech/rivine/modules"
	"github.com/threefoldtech/rivine/types"
)

var (
	rawVersion = "v0.1"
	// Version of the chain binaries.
	//
	// Value is defined by a private build flag,
	// or hardcoded to the latest released tag as fallback.
	Version build.ProtocolVersion
)

const (
	// TokenUnit defines the unit of one Token.
	TokenUnit = "TFB"
	// TokenChainName defines the name of the chain.
	TokenChainName = "tfbchain"
)

// chain names
const (
	NetworkNameStandard = "standard"
	NetworkNameTest     = "testnet"
	NetworkNameDev      = "devnet"
)

// global network config constants
const (
	BlockFrequency types.BlockHeight = 120 // 1 block per 2 minutes on average
)

// GetBlockchainInfo returns the naming and versioning of tfchain.
func GetBlockchainInfo() types.BlockchainInfo {
	return types.BlockchainInfo{
		Name:            TokenChainName,
		NetworkName:     NetworkNameTest,
		CoinUnit:        TokenUnit,
		ChainVersion:    Version,       // use our own blockChain/build version
		ProtocolVersion: build.Version, // use latest available rivine protocol version
	}
}

// GetStandardnetGenesis explicitly sets all the required constants for the genesis block of the standard (prod) net
func GetStandardnetGenesis() types.ChainConstants {
	cfg := types.StandardnetChainConstants()
	return cfg
}

// GetTestnetGenesis explicitly sets all the required constants for the genesis block of the testnet
func GetTestnetGenesis() types.ChainConstants {
	cfg := types.TestnetChainConstants()

	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersionOne
	cfg.GenesisTransactionVersion = types.TransactionVersionOne

	cfg.BlockFrequency = 120

	cfg.MaturityDelay = 720

	// The genesis timestamp
	cfg.GenesisTimestamp = types.Timestamp(1566295200)

	cfg.TargetWindow = 1000

	cfg.MaxAdjustmentUp = big.NewRat(25, 10)
	cfg.MaxAdjustmentDown = big.NewRat(10, 25)

	cfg.FutureThreshold = 120
	cfg.ExtremeFutureThreshold = 600

	cfg.StakeModifierDelay = 2000

	cfg.BlockStakeAging = 64

	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(1.0)

	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(0.1 * 100)

	// allocate block stakes
	cfg.GenesisCoinDistribution = []types.CoinOutput{
		{
			Value:     cfg.CurrencyUnits.OneCoin.Mul64(200 * 1000 * 1000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf"))),
		},
	}

	// allocate block stakes
	cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
		{
			Value:     types.NewCurrency64(2000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf"))),
		},
	}

	return cfg
}

// GetDevnetGenesis explicitly sets all the required constants for the genesis block of the devnet
func GetDevnetGenesis() types.ChainConstants {
	cfg := types.DevnetChainConstants()

	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersionOne
	cfg.GenesisTransactionVersion = types.TransactionVersionOne

	cfg.BlockFrequency = 12

	cfg.MaturityDelay = 10

	// The genesis timestamp
	cfg.GenesisTimestamp = types.Timestamp(1566295200)

	cfg.TargetWindow = 20

	cfg.MaxAdjustmentUp = big.NewRat(120, 100)
	cfg.MaxAdjustmentDown = big.NewRat(100, 120)

	cfg.FutureThreshold = 120
	cfg.ExtremeFutureThreshold = 240

	cfg.StakeModifierDelay = 2000

	cfg.BlockStakeAging = 1024

	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(1.0)

	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(0.1 * 100)

	// allocate block stakes
	cfg.GenesisCoinDistribution = []types.CoinOutput{
		{
			Value: cfg.CurrencyUnits.OneCoin.Mul64(200 * 1000 * 1000),
			// belong to wallet with mnemonic:
			// carbon boss inject cover mountain fetch fiber fit tornado cloth wing dinosaur proof joy intact fabric thumb rebel borrow poet chair network expire else
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
		},
	}

	// allocate block stakes
	cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
		{
			Value: types.NewCurrency64(2000),
			// belong to wallet with mnemonic:
			// carbon boss inject cover mountain fetch fiber fit tornado cloth wing dinosaur proof joy intact fabric thumb rebel borrow poet chair network expire else
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
		},
	}

	return cfg
}

// GetStandardnetBootstrapPeers sets the standard bootstrap node addresses
func GetStandardnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{}
}

// GetTestnetGenesisAuthCoinCondition returns the genesis auth condition used for the testnet

// GetTestnetBootstrapPeers sets the testnet bootstrap node addresses
func GetTestnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
		"bootstrap1.testnet.tfb.threefold.tech:21112",
		"bootstrap2.testnet.tfb.threefold.tech:21112",
		"bootstrap3.testnet.tfb.threefold.tech:21112",
		"bootstrap4.testnet.tfb.threefold.tech:21112",
	}
}

// GetDevnetGenesisAuthCoinCondition returns the genesis auth condition used for the devnet

// GetDevnetBootstrapPeers sets the default devnet bootstrap node addresses
func GetDevnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
		"localhost:21112", // TODO: add port
	}
}

func unlockHashFromHex(hstr string) (uh types.UnlockHash) {
	err := uh.LoadString(hstr)
	if err != nil {
		panic(fmt.Sprintf("func unlockHashFromHex(%s) failed: %v", hstr, err))
	}
	return
}

func init() {
	Version = build.MustParse(rawVersion)
}
