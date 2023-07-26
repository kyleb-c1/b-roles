package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/types"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

var version = "dev"

func main() {
	ctx := context.Background()

	cfg := &Config{}
	cmd, err := cli.NewCmd(ctx, "baton-roles", cfg, validateConfig, getConnector)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	cmd.Version = version

	err = cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func validateConfig(cfg *Config) error {
	// Add any validation logic for your configuration here.
	// Return an error if the configuration is invalid.
	return nil
}

func getConnector(ctx context.Context, cfg *Config) (types.ConnectorServer, error) {
	l := ctxzap.Extract(ctx)
	cb := New(cfg)

	// Print out the roles imported from the config file
	for roleName, _ := range cfg.Roles {
		fmt.Println("Imported role:", roleName)
	}

	c, err := connectorbuilder.NewConnector(ctx, cb)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}

	return c, nil
}
