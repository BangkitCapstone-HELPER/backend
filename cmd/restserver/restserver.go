package restserver

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/spf13/cobra"

	"go.uber.org/fx"
)

var configFile string

//nolint:gochecknoinits
func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c",
		"config/config.yaml", "this parameter is used to start the service application")

	_ = cobra.MarkFlagRequired(pf, "config")
}

var StartCmd = &cobra.Command{
	Use:          "runserver",
	Short:        "Start API server",
	Example:      "{execfile} server -c config/config.yaml",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(app.Module).Run()
	},
}
