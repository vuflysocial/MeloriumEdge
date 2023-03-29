package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "meloriumedge/testutil/keeper"
	"meloriumedge/x/meloriumedge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MeloriumedgeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
