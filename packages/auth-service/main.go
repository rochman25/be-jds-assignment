package main

import (
	"auth-service/database/migration"
	config "auth-service/pkg"
	"auth-service/src/factory"
	"auth-service/src/http"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	ctx := context.Background()
	config.LoadEnv("../../.env")
	rootCmd := &cobra.Command{}

	// define rest command
	restCmd := cobra.Command{
		Use:   "rest",
		Short: "Rest is a command to start rest server",
		Run: func(cmd *cobra.Command, args []string) {
			f := factory.NewFactory(ctx).BuildRestFactory()

			g := gin.New()
			http.NewHttp(g, f)

			if err := g.Run(fmt.Sprintf(":%d", config.AppPort())); err != nil {
				log.Fatal("Can't start server.")
			}
		},
	}

	// define migration command
	migrationCmd := cobra.Command{
		Use:   "migration",
		Short: "Migration is a command to migrate database",
	}

	// define migration init command
	migrationInit := cobra.Command{
		Use:   "init",
		Short: "Init is a command to create migration table",
		Run: func(cmd *cobra.Command, args []string) {
			migration.FirstMigrate()
		},
	}

	// define migration creat command
	migrationCreate := cobra.Command{
		Use:   "create",
		Short: "Create is a command to create migration file",
		Long:  "Example: go run main.go migration file create_table_user",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("migration file name is required")
			}

			file := args[0]

			if err := migration.CreateMigrationFile(file); err != nil {
				log.Fatalf("failed to create migration file: %v", err)
			}
		},
	}

	// define migration file command
	migrationFile := cobra.Command{
		Use:   "file",
		Short: "File is a command to migrate file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("migration file name is required")
			}

			file := args[0]

			if err := migration.Migrate(file); err != nil {
				log.Fatalf("failed to migrate file: %v", err)
			}
		},
	}

	// register root command
	rootCmd.AddCommand(
		&restCmd,
		&migrationCmd,
	)

	// register migration command
	migrationCmd.AddCommand(
		&migrationInit,
		&migrationCreate,
		&migrationFile,
	)

	// execute root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root command: %v", err)
	}
}
