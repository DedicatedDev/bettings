package bettings_test

import (
	"testing"

	keepertest "github.com/DedicatedDev/bettings/testutil/keeper"
	"github.com/DedicatedDev/bettings/testutil/nullify"
	"github.com/DedicatedDev/bettings/x/bettings"
	"github.com/DedicatedDev/bettings/x/bettings/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BettingsKeeper(t)
	bettings.InitGenesis(ctx, *k, genesisState)
	got := bettings.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
