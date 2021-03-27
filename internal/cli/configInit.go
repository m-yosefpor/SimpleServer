package cli

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Spf struct {
	C   *cobra.Command
	V   *viper.Viper
	Cfg *Config
}

func InitCli() *Spf {
	s := &Spf{
		V:   viper.New(),
		Cfg: NewConfig(),
	}
	s.C = NewRootCommand(s.Cfg)
	return s
}

func (s *Spf) ReadConfig(name string) error {
	s.V.SetConfigName(name)
	s.V.SetConfigType("yaml")
	s.V.AddConfigPath(".")
	home, err := homedir.Dir() // side-effect: Find home directory.
	if err != nil {
		return err
	}
	s.V.AddConfigPath(home)

	s.V.AutomaticEnv() // side-effect: read in environment variables that match

	if err := s.V.ReadInConfig(); err != nil { // side-effect: read config file
		return err
		// It's okay if there isn't a config file
		// if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// return err
		// }
	}
	if err := s.V.Unmarshal(s.Cfg); err != nil {
		return err
	}
	log.Info(s.Cfg.Listen.Ip)
	log.Info(s.Cfg.Listen.Port)
	log.Info(s.Cfg.Auth.Token)
	return nil
}
