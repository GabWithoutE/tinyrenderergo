package tinyrendercli

import (
	"github.com/spf13/cobra"
)

type TinyRendererCLIApp struct {
	command *cobra.Command
}

func NewTinyRenderCLIApp() (*TinyRendererCLIApp, error) {
	return &TinyRendererCLIApp{

	}, nil
}