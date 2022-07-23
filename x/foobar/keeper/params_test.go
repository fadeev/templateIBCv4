package keeper_test

import (
	"testing"

	testkeeper "github.com/fadeev/templateIBCv4/testutil/keeper"
	"github.com/fadeev/templateIBCv4/x/foobar/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FoobarKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
