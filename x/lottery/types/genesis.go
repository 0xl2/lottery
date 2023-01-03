package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LotteryInfo:       nil,
		StoredLotteryList: []StoredLottery{},
		StoredBetList:     []StoredBet{},
		BetInfo:           nil,
		LotteryData:       nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedLottery
	storedLotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredLotteryList {
		index := string(StoredLotteryKey(elem.Index))
		if _, ok := storedLotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedLottery")
		}
		storedLotteryIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedBet
	storedBetIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredBetList {
		index := string(StoredBetKey(elem.Index))
		if _, ok := storedBetIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedBet")
		}
		storedBetIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
