package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/compose"
	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/config"
	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/database"
	"github.com/shadracnicholas/home-automation/tools/bolt/pkg/service"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/output"
)

var (
	dbSchemaCmd = &cobra.Command{
		Use:   "schema [service.foo] [service.bar]",
		Short: "apply a service's schema",
		Long:  "Apply one or more services' schemas. If run without arguments, all service schemas are applied.",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := compose.New()
			if err != nil {
				output.Fatal("Failed to init compose: %v", err)
			}

			all, err := cmd.Flags().GetBool("all")
			if err != nil {
				output.Fatal("Failed to parse all flag: %v", err)
			}

			services := service.Expand(args)

			if all {
				var err error
				services, err = c.ListAll()
				if err != nil {
					output.Fatal("Failed to list all services: %v", err)
				}
			}

			db := config.Get().Database

			for _, serviceName := range services {
				schema, err := database.GetDefaultSchema(serviceName)
				if err != nil {
					output.Fatal("Failed to get schema for %s: %v", serviceName, err)
				}

				// Silently skip services that don't have a schema
				if schema == "" {
					continue
				}

				if err := database.New(c, &db).ApplySchema(serviceName, schema); err != nil {
					output.Fatal("Failed to apply schema: %v", err)
				}
			}
		},
	}
)

func init() {
	dbCmd.AddCommand(dbSchemaCmd)
	dbSchemaCmd.Flags().Bool("all", false, "apply all schemas")
}
