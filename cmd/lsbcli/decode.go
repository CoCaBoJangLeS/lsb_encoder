package lsbcli

import (
	"github.com/spf13/cobra"
)

// DecodeCommand decodes a message hidden inside an image
var DecodeCommand = &cobra.Command{
	Use:   "decode",
	Short: "Decode image",
	Long:  "Decodes a message hidden inside an image",
	RunE:  decodeCommandFunction,
}

func decodeCommandFunction(command *cobra.Command, args []string) (err error) {

}
