package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/gridchain/gridiron/v11/app"
	gridirond "github.com/gridchain/gridiron/v11/cmd/gridirond"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := gridirond.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"gridiron-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "gridiron_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, "gridirond", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := gridirond.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "GRIDIROND", app.DefaultNodeHome)
	require.Error(t, err)
}
