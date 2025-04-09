package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}

type CobraConfig struct {
	Use   string
	Short string
	Long  string
}

func CobraInit(config CobraConfig) *cobra.Command {
	return &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
	}
}
