package meloriumedge_test

import (
	"testing"

	keepertest "meloriumedge/testutil/keeper"
	"meloriumedge/testutil/nullify"
	"meloriumedge/x/meloriumedge"
	"meloriumedge/x/meloriumedge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MeloriumedgeKeeper(t)
	meloriumedge.InitGenesis(ctx, *k, genesisState)
	got := meloriumedge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
