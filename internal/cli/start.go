package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newStartCmd(c *Config) *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start the webserver",
		Run: func(cmd *cobra.Command, args []string) {
			start(c)
		},
	}
	startCmd.Flags().IntVarP(&c.Listen.Port, "port", "p", 1, "port to bind to")
	if err := viper.BindPFlag("port", startCmd.Flags().Lookup("port")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
	startCmd.Flags().IntVarP(&c.Verbose, "verbose", "v", 2, "verbosity")
	if err := viper.BindPFlag("verbose", startCmd.Flags().Lookup("verbose")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
	return startCmd
}

func start(c *Config) {

	log.Info("Log level: ", c.Verbose)
	log.Info(c.Listen.Port)
	log.Warn("I'm here")

}
