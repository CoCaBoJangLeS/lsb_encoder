package lsbcli

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"lsb_encoder/pkg/process"
// 	"os"
// 	"path/filepath"
// )

// func parseFlags() (*process.Flags, []error) {
// 	flag.Parse()

// 	var msgFilePath string
// 	var errs []error
// 	srcFilePath, err := filepath.Abs(*srcFile)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}
// 	outFilePath, err := filepath.Abs(*outDir)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}
// 	if *msgFile != "" {
// 		msgFilePath, err = filepath.Abs(*msgFile)
// 		if err != nil {
// 			errs = append(errs, err)
// 		}
// 	}
// 	// If attempting to encode without a source message
// 	if !*decode && (msgFilePath == "") && (*text == "") {
// 		errs = append(errs, fmt.Errorf("need a text source to encode"))
// 	}
// 	return &process.Flags{
// 		SrcFile:     srcFilePath,
// 		OutputDir:   outFilePath,
// 		Text:        *text,
// 		MessageFile: msgFilePath,
// 		BitOpt:      *bitOpt,
// 		Decode:      *decode,
// 		Rot13:       *rot13,
// 		Base16:      *base16,
// 		Base32:      *base32,
// 		Base64:      *base64,
// 		Base85:      *base85,
// 		Complex:     *complex,
// 	}, errs
// }

// func main() {
// 	flags, errs := parseFlags()
// 	if len(errs) != 0 {
// 		for _, err := range errs {
// 			log.Println(err)
// 		}
// 		os.Exit(1)
// 	}
// 	if flags.Decode {
// 		extractSecret, err := process.ParseExtractSecret(flags)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = process.Extract(extractSecret)
// 		if err != nil {
// 			panic(err)
// 		}
// 		os.Exit(0)
// 	}
// 	embedSecret, err := process.ParseEmbedSecret(flags)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	err = process.Embed(embedSecret)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// }
