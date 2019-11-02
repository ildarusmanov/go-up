package commands

import (
	"github.com/spf13/cobra"
)

// NewRootCmd() creates new root command
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "run",
		Short: "run goup command",
		Long:  ``,
	}

	initCmd, initCmdInit := NewInitCommand()

	// run commands initializers
	initCmdInit()

	// add commands
	rootCmd.AddCommand(initCmd)

	return rootCmd
}

// Execute executes the root command.
func Execute() error {
	return NewRootCmd().Execute()
}
