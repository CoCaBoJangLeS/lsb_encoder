package process

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
)

// Embed embeds the Secret's data into the source file and writes it to the output directory
func Embed(secret *Secret) (string, error) {
	var loadedImage image.Image
	var format string
	var err error
	// Read the Source file
	sourceFile, err := os.Open(secret.SourcePath)
	if err != nil {
		return "", fmt.Errorf("error reading Source file: (%v)", err)
	}
	defer sourceFile.Close()
	_, format, err = image.Decode(sourceFile)
	if err != nil {
		return "", fmt.Errorf("error decoding source file: (%v)", err)
	}
	// Reset the file's reader to beginning
	sourceFile.Seek(0, 0)

	// Do all the work, can totally be cleaned up & refactored
	// cause it's kinda a mess right now
	if format == "png" {
		// =====
		// Handle PNG
		// ====
		loadedImage, err = png.Decode(sourceFile)
		if err != nil {
			return "", fmt.Errorf("error decoding PNG file: (%v)", err)
		}
		embedded, err := EmbedMsgInImage(secret, loadedImage)
		if err != nil {
			return "", fmt.Errorf("error embedding message in file: (%v)", err)
		}
		newFile, err := os.Create(filepath.Join(secret.OutputDir, "output."+format))
		if err != nil {
			return "", fmt.Errorf("error creating output file: (%v)", err)
		}
		defer newFile.Close()
		err = png.Encode(newFile, embedded)
		if err != nil {
			return "", fmt.Errorf("error encoding new PNG image: (%v)", err)
		}
		return filepath.Join(secret.OutputDir, "output."+format), nil
	} else if format == "jpeg" {
		// =====
		// Handle JPEG
		// ====
		loadedImage, err = jpeg.Decode(sourceFile)
		if err != nil {
			return "", fmt.Errorf("error decoding JPEG file: (%v)", err)
		}
		embedded, err := EmbedMsgInImage(secret, loadedImage)
		if err != nil {
			return "", fmt.Errorf("error embedding message in file: (%v)", err)
		}
		newFile, err := os.Create(filepath.Join(secret.OutputDir, "output_"+format+".png"))
		if err != nil {
			return "", fmt.Errorf("error creating output file: (%v)", err)
		}
		defer newFile.Close()
		err = png.Encode(newFile, embedded)
		if err != nil {
			return "", fmt.Errorf("error encoding new JPEG image: (%v)", err)
		}
		return filepath.Join(secret.OutputDir, "output_"+format+".png"), nil
	} else if format == "bmp" {
		// =====
		// Handle BMP
		// ====
		loadedImage, err = bmp.Decode(sourceFile)
		if err != nil {
			return "", fmt.Errorf("error decoding BMP file: (%v)", err)
		}
		embedded, err := EmbedMsgInImage(secret, loadedImage)
		if err != nil {
			return "", fmt.Errorf("error embedding message in file: (%v)", err)
		}
		newFile, err := os.Create(filepath.Join(secret.OutputDir, "output."+format))
		if err != nil {
			return "", fmt.Errorf("error creating output file: (%v)", err)
		}
		defer newFile.Close()
		err = bmp.Encode(newFile, embedded)
		if err != nil {
			return "", fmt.Errorf("error encoding new JPEG image: (%v)", err)
		}
		return filepath.Join(secret.OutputDir, "output."+format), nil
	} else if format == "gif" {
		// =====
		// Handle GIF
		// ====
		loadedGIF, err := gif.DecodeAll(sourceFile)
		if err != nil {
			return "", fmt.Errorf("error decoding GIF file: (%v)", err)
		}
		embedded, err := EmbedMsgInGIF(secret, loadedGIF)
		if err != nil {
			return "", fmt.Errorf("error embedding message in file: (%v)", err)
		}
		newFile, err := os.Create(filepath.Join(secret.OutputDir, "output."+format))
		if err != nil {
			return "", fmt.Errorf("error creating output file: (%v)", err)
		}
		defer newFile.Close()
		err = gif.EncodeAll(newFile, embedded)
		if err != nil {
			return "", err
		}
	} else {
		// =====
		// Handle ???
		// ====
		return "", fmt.Errorf("unsupported source file format: %v", format)
	}
	return "", nil
}

// Extract extracts the Secret's data and writes it to the output directory
func Extract(secret *Secret) error {
	var loadedImage image.Image
	// Read the Source file
	sourceFile, err := os.Open(secret.SourcePath)
	if err != nil {
		return fmt.Errorf("error reading Source file: (%v)", err)
	}
	defer sourceFile.Close()
	_, format, err := image.Decode(sourceFile)
	if err != nil {
		return fmt.Errorf("error decoding source file: (%v)", err)
	}
	// Reset the file's reader to beginning
	sourceFile.Seek(0, 0)
	if format == "png" {
		// =====
		// Handle PNG
		// ====
		loadedImage, err = png.Decode(sourceFile)
		if err != nil {
			return fmt.Errorf("error decoding PNG file: (%v)", err)
		}
		extracted, err := ExtractMsgFromImage(secret, loadedImage)
		if err != nil {
			return fmt.Errorf("error extracting message from file: (%v)", err)
		}
		err = WriteFile(extracted.Message, extracted.OutputDir, extracted.DataHeader.Type)
		if err != nil {
			return fmt.Errorf("error creating output file: (%v)", err)
		}
		return nil
	} else if format == "jpeg" {
		// =====
		// Handle JPEG
		// ====
		loadedImage, err = jpeg.Decode(sourceFile)
		if err != nil {
			return fmt.Errorf("error decoding JPEG file: (%v)", err)
		}
		extracted, err := ExtractMsgFromImage(secret, loadedImage)
		if err != nil {
			return fmt.Errorf("error extracting message from file: (%v)", err)
		}
		err = WriteFile(extracted.Message, extracted.OutputDir, extracted.DataHeader.Type)
		if err != nil {
			return fmt.Errorf("error creating output file: (%v)", err)
		}
		return nil
	} else if format == "bmp" {
		// =====
		// Handle BMP
		// ====
		loadedImage, err = bmp.Decode(sourceFile)
		if err != nil {
			return fmt.Errorf("error decoding BMP file: (%v)", err)
		}
		extracted, err := ExtractMsgFromImage(secret, loadedImage)
		if err != nil {
			return fmt.Errorf("error extracting message from file: (%v)", err)
		}
		err = WriteFile(extracted.Message, extracted.OutputDir, extracted.DataHeader.Type)
		if err != nil {
			return fmt.Errorf("error creating output file: (%v)", err)
		}
		return nil
	} else if format == "gif" {
		// =====
		// Handle GIF
		// ====
		loadedGIF, err := gif.DecodeAll(sourceFile)
		if err != nil {
			return fmt.Errorf("error decoding GIF file: (%v)", err)
		}
		extracted, err := ExtractMsgFromGif(secret, loadedGIF)
		if err != nil {
			return fmt.Errorf("error extracting message from file: (%v)", err)
		}
		err = WriteFile(extracted.Message, extracted.OutputDir, extracted.DataHeader.Type)
		if err != nil {
			return fmt.Errorf("error creating output file: (%v)", err)
		}
		return nil
	}
	// =====
	// Handle ???
	// ====
	return fmt.Errorf("unsupported source file format: %v", format)
}
