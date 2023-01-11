package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Name defines the application name of the lottery network.
	Name = "lott"

	// BondDenom defines the native staking token denomination.
	BondDenom = "lott"

	// DisplayDenom defines the name, symbol, and display value of the lott token.
	DisplayDenom = "LOTT"

	// DefaultGasLimit - set to the same value as cosmos-sdk flags.DefaultGasLimit
	// this value is currently only used in tests.
	DefaultGasLimit = 200000
)

var (
	// ProtocolMinGasPrice is a consensus controlled gas price. Each validator must set his
	// `minimum-gas-prices` in app.toml config to value above ProtocolMinGasPrice.
	// Transactions with gas-price smaller than ProtocolMinGasPrice will fail during DeliverTx.
	ProtocolMinGasPrice = sdk.NewDecCoinFromDec(BondDenom, sdk.MustNewDecFromStr("0.00"))
)
