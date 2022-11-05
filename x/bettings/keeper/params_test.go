package keeper_test

import (
	"testing"

	testkeeper "github.com/DedicatedDev/bettings/testutil/keeper"
	"github.com/DedicatedDev/bettings/x/bettings/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BettingsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
