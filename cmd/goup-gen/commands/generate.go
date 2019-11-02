package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func NewGenerateCommand() (*cobra.Command, func()) {
	return &cobra.Command{
		Use:   "ge [string to echo]",
		Short: "Echo anything to the screen",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   runGenerateCmd,
	}, initGenerateCmd
}

func runGenerateCmd(cmd *cobra.Command, args []string) {
	log.Printf("%+v", args)
}

func initGenerateCmd() {

}
