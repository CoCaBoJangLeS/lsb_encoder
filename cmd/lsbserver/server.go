package lsbserver

import (
	"lsb_encoder/pkg/api"

	"github.com/spf13/cobra"
)

// ServerCommand starts the HTTP Server for handling API requests and the WebApp
var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Run HTTP server",
	Long:  "Starts an HTTP Server for the WebApp and API calls",
	RunE:  serverCommandFunction,
}

func serverCommandFunction(command *cobra.Command, args []string) (err error) {
	s := api.NewServer("8080")
	s.Logger.Infof("Server starting, listening on http://localhost:8080")
	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
	return nil
}
