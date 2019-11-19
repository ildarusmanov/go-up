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
	generateCmd, initGenerateCmd := NewGenerateCommand()
	installCmd, initInstallCmd := NewInstallCommand()

	// run commands initializers
	initCmdInit()
	initGenerateCmd()
	initInstallCmd()

	// add commands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(installCmd)

	return rootCmd
}

// Execute executes the root command.
func Execute() error {
	return NewRootCmd().Execute()
}
