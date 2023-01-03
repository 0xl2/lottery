package types_test

import (
	"testing"

	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
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

				LotteryInfo: &types.LotteryInfo{
					NextId: 35,
				},
				StoredLotteryList: []types.StoredLottery{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				StoredBetList: []types.StoredBet{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				BetInfo: &types.BetInfo{
					BetId: 64,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedLottery",
			genState: &types.GenesisState{
				StoredLotteryList: []types.StoredLottery{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedBet",
			genState: &types.GenesisState{
				StoredBetList: []types.StoredBet{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
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
