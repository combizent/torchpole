package torchpole

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var CmdRun = &cobra.Command{
	Use: "run",

	Short: "Torchpole run cmd.",
	Long: `Torchpole run cmd.

Find more torchpole information at: https://github.com/combizent/torchpole`,

	RunE: func(cmd *cobra.Command, args []string) error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		fmt.Println("Shutting down server ...")
		// ...
		return nil
	},
}
