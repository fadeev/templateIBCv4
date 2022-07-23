package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fadeev/templateIBCv4/testutil/keeper"
	"github.com/fadeev/templateIBCv4/x/foobar/keeper"
	"github.com/fadeev/templateIBCv4/x/foobar/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FoobarKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
