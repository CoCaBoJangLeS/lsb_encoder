package lsbcli

import (
	"github.com/spf13/cobra"
)

var (
	text, inFile, complex   *string
	bitOpt                  *int
	r13, b16, b32, b64, b85 *bool
)

// EncodeCommand hides a message inside an image
var EncodeCommand = &cobra.Command{
	Use:   "encode",
	Short: "Encode image",
	Long:  "Encodes a message to be hidden inside an image",
	RunE:  encodeCommandFunction,
}

func encodeCommandFunction(command *cobra.Command, args []string) (err error) {
	command.Flags().StringVarP(text, "text", "t", "", "The text to be embedded")
	command.Flags().StringVarP(inFile, "infile", "i", "", "The input file to be embedded")
	command.Flags().StringVarP(complex, "complex", "c", "", "A comma separated list of encoding types, applied in order (limit 5)")
	command.Flags().IntVarP(bitOpt, "bits", "b", 2, "1 = Last Bit Only, 2 = 2-3-3/R-G-B (Default 2)")
	command.Flags().BoolVarP(r13, "rot13", "r13", false, "Apply r13 encoding to the message before encoding")
	command.Flags().BoolVarP(b16, "base16", "b16", false, "Apply b16 encoding to the message before encoding")
	command.Flags().BoolVarP(b32, "base32", "b32", false, "Apply b32 encoding to the message before encoding")
	command.Flags().BoolVarP(b64, "base64", "b64", false, "Apply b64 encoding to the message before encoding")
	command.Flags().BoolVarP(b85, "base85", "b85", false, "Apply b85 encoding to the message before encoding")

}
