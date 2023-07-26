package main

import (
	"flag"
	"fmt"

	"github.com/ConductorOne/baton-sdk/pkg/client"
	"github.com/ConductorOne/baton-sdk/pkg/connector"
	"github.com/yourusername/baton-roles/pkg/config"
)

func main() {
	clientID := flag.String("client_id", "", "Client ID")
	clientSecret := flag.String("client_secret", "", "Client Secret")
	configFile := flag.String("config", "", "Path to the YAML configuration file")

	flag.Parse()

	if *clientID == "" || *clientSecret == "" || *configFile == "" {
		fmt.Println("You must provide a client_id, client_secret, and config file.")
		return
	}

	cfg, err := config.Load(*configFile)
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	cli, err := client.New(*clientID, *clientSecret)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}

	conn := connector.New(cli, cfg)
	conn.Start()
}
