package cmd

import (
	"github.com/spf13/viper"
)

type TinyRenderCLIConfig struct {
	OutputFilename string `mapstructure:"outputfilename"`
	FillMethod string `mapstructure:"fillmethod"`
}

func LoadConfig() (*TinyRenderCLIConfig, error) {
	v := viper.New()
	v.AddConfigPath()
}