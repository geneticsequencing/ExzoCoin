package root

import (
	"fmt"
	"os"

	"github.com/ExzoNetwork/ExzoCoin/command/backup"
	"github.com/ExzoNetwork/ExzoCoin/command/genesis"
	"github.com/ExzoNetwork/ExzoCoin/command/helper"
	"github.com/ExzoNetwork/ExzoCoin/command/ibft"
	"github.com/ExzoNetwork/ExzoCoin/command/license"
	"github.com/ExzoNetwork/ExzoCoin/command/monitor"
	"github.com/ExzoNetwork/ExzoCoin/command/peers"
	"github.com/ExzoNetwork/ExzoCoin/command/secrets"
	"github.com/ExzoNetwork/ExzoCoin/command/server"
	"github.com/ExzoNetwork/ExzoCoin/command/status"
	"github.com/ExzoNetwork/ExzoCoin/command/txpool"
	"github.com/ExzoNetwork/ExzoCoin/command/version"
	"github.com/ExzoNetwork/ExzoCoin/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Exzo Network is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
