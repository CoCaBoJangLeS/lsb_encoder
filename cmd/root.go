package main

import (
	"fmt"
	"os"

	"lsb_encoder/cmd/lsbserver"

	"github.com/spf13/cobra"
)

var (
	srcFile, outDir string
	help            bool
	rootCommand     = &cobra.Command{
		Use:   "server",
		Short: "",
		Long:  "",
	}
)

func init() {
	rootCommand.PersistentFlags().StringVarP(&srcFile, "srcfile", "s", "", "Path to the source file")
	rootCommand.PersistentFlags().StringVarP(&outDir, "outdir", "o", "", "Directory to store the output")
	rootCommand.PersistentFlags().BoolVarP(&help, "help", "h", false, "Print out help text with examples")
	rootCommand.AddCommand(lsbserver.ServerCommand)
	// rootCommand.AddCommand(lsbcli.EncodeCommand)
	// rootCommand.AddCommand(lsbcli.DecodeCommand)
}

// Execute runs the main rootCommand
func Execute() {
	if help {
		fmt.Println(`
		Flag Options:
		- s # or --srcfile /path/to/input/source.png (.gif, .bmp, or .jpeg)
		- o # or --outdir /path/to/output/ Directory to save output.png (.gif, .bmp, or .jpeg)
		- t # or --text "The Secret Message to embed"
		- m # or --msgfile /path/to/secret_message.txt (can be anything)
		- i # or --stdin The secret message to embed comes from stdin (ex: pipe command)
		- b # or --bitopt Option for which LSBs to embed: 1 = Last Bit Only, 2 = 2-3-3/R-G-B method (default)
		- d # or --decode Extract a message from an already embedded file
		- r13 # or --rot13 Apply Rot13 pre encoding to the message before embedding
		- b16 # or --base16 Apply Base16 pre encoding to the message before embedding
		- b32 # or --base32 Apply Base32 pre encoding to the message before embedding
		- b64 # or --base64 Apply Base64 pre encoding to the message before embedding
		- b85 # or --base85 Apply Base85 pre encoding to the message before embedding
		- c # or --complex A Comma separated list(no space) of encoding types, applied in the order they appear (limit 5)
		- h # or --help Print out help text

		# Example commands
		# Simple
		go run ./cmd/lsb_encoder/ \
			--srcfile ~/Desktop/Pics/kitty_cat.jpeg \
			--outdir ~/Desktop/Pics -base64 \
			--text "Kitty Cat"

		go run ./cmd/lsb_encoder/ --decode --b64 \
			-s ~/Desktop/Pics/output_jpeg.png \
			-o ~/Desktop/Pics \

		# Fancy
  	# embed a message in a small image file, like my_avatar.png
  	go run ./cmd/lsb_encoder/ \
    	-s ~/Desktop/Pics/my_avatar.png \
    	-o ~/Desktop/Pics/Output \
    	--text "Shhhh, don't tell anyone this is hidden in my avatar."

  	# embed the output from above in a wallpaper
  	go run ./cmd/lsb_encoder/ \
    	-s ~/Desktop/Pics/really_cool_wallpaper.jpeg \
    	-o ~/Desktop/Pics/Output \
    	--msgfile ~/Desktop/Pics/Output/output.png
		`)
		os.Exit(0)
	}
	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}

func main() {
	rootCommand.Execute()
}
