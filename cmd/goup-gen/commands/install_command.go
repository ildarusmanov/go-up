package commands

import (
	"log"
	"os"

	"github.com/ildarusmanov/go-up/internal/config"
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
	return func() {

	}
}

// runInstallCmd runs Init command
func runInstallCmd(cmd *cobra.Command, args []string) {
	os.Mkdir(config.DefaultGoupDir, os.FileMode(0777))
	os.Mkdir(config.DefaultTemplatesDir, os.FileMode(0777))
	f, err := os.Create(config.DefaultTemplatesConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}
