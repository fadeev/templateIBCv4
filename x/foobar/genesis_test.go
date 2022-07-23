package foobar_test

import (
	"testing"

	keepertest "github.com/fadeev/templateIBCv4/testutil/keeper"
	"github.com/fadeev/templateIBCv4/testutil/nullify"
	"github.com/fadeev/templateIBCv4/x/foobar"
	"github.com/fadeev/templateIBCv4/x/foobar/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FoobarKeeper(t)
	foobar.InitGenesis(ctx, *k, genesisState)
	got := foobar.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
