package keeper_test

import (
	"context"
	"testing"

	keepertest "belkyc/testutil/keeper"
	"belkyc/x/test/keeper"
	"belkyc/x/test/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
