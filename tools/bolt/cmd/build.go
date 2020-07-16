package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/compose"
	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/service"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/output"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build [service.foo] [service.bar]...",
		Short: "build a service",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			services := service.Expand(args)

			c, err := compose.New()
			if err != nil {
				output.Fatal("Failed to init compose: %v", err)
			}

			if err := c.Build(services); err != nil {
				output.Fatal("Failed to build: %v", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}
