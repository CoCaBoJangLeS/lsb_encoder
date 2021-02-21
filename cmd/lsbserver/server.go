package lsbserver

import (
	"net/http"

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
	server := &http.Server{
		Addr:    "",
		Handler: nil,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	return err
}
