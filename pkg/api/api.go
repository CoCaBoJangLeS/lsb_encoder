package api

import (
	"fmt"
	"io"
	"lsb_encoder/pkg/process"
	"net/http"
	"os"

	"github.com/gorilla/schema"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Server wraps http.Server and adds a zap.Logger
type Server struct {
	Server *http.Server
	Logger *zap.SugaredLogger
}

type endpoint struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

// Embed stores input data from the UI
type Embed struct {
	Message   string `schema:"message"`
	PreEncode string `schema:"pre_encoding"`
}

func endpoints(logger *zap.SugaredLogger) []endpoint {
	return []endpoint{
		{
			Path:        "/api/embed",
			Method:      http.MethodPost,
			HandlerFunc: handleEmbedMessage(logger),
		},
		{
			Path:        "/api/extract",
			Method:      http.MethodPost,
			HandlerFunc: handleExtractMessage(logger),
		},
	}
}

// NewServer intializes a new server hosting our endpoints
func NewServer(port string) *Server {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	endpoints := endpoints(sugar)
	router := mux.NewRouter().StrictSlash(true)
	for _, endpoint := range endpoints {
		router.HandleFunc(endpoint.Path, endpoint.HandlerFunc)
	}
	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf("localhost:%s", port),
			Handler: router,
		},
		Logger: sugar,
	}
}

func handleEmbedMessage(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("handling embed message")
		r.ParseMultipartForm(32 << 20) // 32 Mb
		upload, header, err := r.FormFile("file")
		if err != nil {
			logger.Errorf("err parsing file upload: %v", err)
		}
		defer upload.Close()

		tempFilename := fmt.Sprintf("./%s.tmp", header.Filename)
		tempFile, err := os.OpenFile(tempFilename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			logger.Errorf("err creating temporary file: %v", err)
		}
		defer func() {
			os.Remove(tempFilename)
		}()

		_, err = io.Copy(tempFile, upload)
		if err != nil {
			logger.Errorf("err saving temporary file: %v", err)
		}
		tempFile.Close()

		input := &Embed{}

		err = schema.NewDecoder().Decode(input, r.Form)
		if err != nil {
			logger.Errorf("err decoding form data: %v", err)
		}
		flags := &process.Flags{
			SrcFile:   tempFilename,
			OutputDir: "./",
			Text:      input.Message,
		}
		switch input.PreEncode {
		case "r13":
			flags.Rot13 = true
		case "b16":
			flags.Base16 = true
		case "b32":
			flags.Base32 = true
		case "b64":
			flags.Base64 = true
		case "b85":
			flags.Base85 = true
		default:
			break
		}
		secret, err := process.ParseEmbedSecret(flags)
		if err != nil {
			logger.Errorf("err parsing embed request: %v", err)
		}
		outpath, err := process.Embed(secret)
		if err != nil {
			logger.Errorf("err embedding secret: %v", err)
		}
		output, err := os.Open(outpath)
		if err != nil {
			logger.Errorf("err opening output file: %v", err)
		}
		_, err = io.Copy(w, output)
		if err != nil {
			logger.Errorf("err writing output to response: %v", err)
		}
		logger.Infof("finished embed message")
	}
}

func handleExtractMessage(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("handling extract message")
	}
}
