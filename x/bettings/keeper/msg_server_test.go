package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DedicatedDev/bettings/testutil/keeper"
	"github.com/DedicatedDev/bettings/x/bettings/keeper"
	"github.com/DedicatedDev/bettings/x/bettings/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BettingsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
