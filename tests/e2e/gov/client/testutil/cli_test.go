//go:build e2e
// +build e2e

package testutil

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/testutil"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 1
	suite.Run(t, testutil.NewIntegrationTestSuite(cfg))

	dp := v1.NewDepositParams(sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, v1.DefaultMinDepositTokens)), time.Duration(15)*time.Second)
	vp := v1.NewVotingParams(time.Duration(5) * time.Second)
	genesisState := v1.DefaultGenesisState()
	genesisState.DepositParams = &dp
	genesisState.VotingParams = &vp
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["gov"] = bz
	suite.Run(t, testutil.NewDepositTestSuite(cfg))
}
