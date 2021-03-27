package cli

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitializeConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			log.Println("config specified but unable to read it, using defaults")
		}
	}
}

func Read(name string) error {
	viper.SetConfigName(name)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("cannot read config file: %w", err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	return nil
}
