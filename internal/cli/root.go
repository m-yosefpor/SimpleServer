package cli

import (
	"github.com/m-yosefpor/simpleserver/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {

	cfg := &config.Config{}
	// Define our command
	rootCmd := &cobra.Command{
		Use:   "simpleserver",
		Short: "A simple filesharing webserver",
		Long:  `A simple webserver to download and upload files intended for LAN.`,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return config.InitializeConfig(cfg)
		},
	}
	rootCmd.AddCommand(newVersionCmd(cfg))
	rootCmd.AddCommand(newStartCmd(cfg))

	return rootCmd
}
