package rootcmd

import (
	"errors"
	"os"

	"github.com/BangkitCapstone-HELPER/backend/cmd/restserver"
	"github.com/spf13/cobra"
)

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(restserver.StartCmd)
}

var rootCmd = &cobra.Command{
	Use:          "pillar-policy",
	Short:        "pillar-policy",
	SilenceUsage: true,
	Long:         `pillar-policy`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(
				"requires at least one arg, " +
					"you can view the available parameters through `--help`",
			)
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run:               func(cmd *cobra.Command, args []string) {},
}

var osExit = os.Exit

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		osExit(-1)
	}
}
