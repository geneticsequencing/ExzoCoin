package version

import (
	"github.com/ExzoNetwork/ExzoCoin/command"
	"github.com/ExzoNetwork/ExzoCoin/versioning"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Returns the current Exzo Network version",
		Args:  cobra.NoArgs,
		Run:   runCommand,
	}
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	outputter.SetCommandResult(
		&VersionResult{
			Version:   versioning.Version,
			Commit:    versioning.Commit,
			Branch:    versioning.Branch,
			BuildTime: versioning.BuildTime,
		},
	)
}
