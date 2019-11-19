package commands

import (
	"github.com/spf13/cobra"
)

// NewInstallCommand initializes Init command
func NewInstallCommand() (*cobra.Command, func()) {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install go-up",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run:   runInstallCmd,
	}

	return cmd, initInstallCmd(cmd)
}

// initInstallCmd initializes Init command
func initInstallCmd(cmd *cobra.Command) func() {
	return func() {}
}

// runInstallCmd runs Init command
func runInstallCmd(cmd *cobra.Command, args []string) {
}
