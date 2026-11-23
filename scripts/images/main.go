package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

const (
	maxWidth              = 200
	preferredMinWidth     = 100
	targetBytes           = 150 * 1024
	maxBytes              = 250 * 1024
	maxSourcePixels       = 25_000_000
	maxKernelScratchBytes = 16 * 1024 * 1024
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run scripts/images/main.go INPUT_DIRECTORY NEW_OUTPUT_DIRECTORY")
		os.Exit(2)
	}
	if err := optimizeDirectory(os.Args[1], os.Args[2], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func optimizeDirectory(input, output string, report io.Writer) (err error) {
	entries, err := os.ReadDir(input)
	if err != nil {
		return err
	}
	var sources []string
	for _, entry := range entries {
		switch filepath.Ext(entry.Name()) {
		case ".jpg", ".png":
			if !entry.Type().IsRegular() {
				return fmt.Errorf("%s: expected a regular file", entry.Name())
			}
			sources = append(sources, entry.Name())
		}
	}
	if len(sources) == 0 {
		return fmt.Errorf("no .jpg or .png images in %s", input)
	}
	if err := os.Mkdir(output, 0755); err != nil {
		return fmt.Errorf("create new output directory: %w", err)
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, os.RemoveAll(output))
		}
	}()
	for _, name := range sources {
		if err := optimizeImage(filepath.Join(input, name), filepath.Join(output, name), report); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}
	return nil
}

func optimizeImage(source, destination string, report io.Writer) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	config, format, err := image.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("read image header: %w", err)
	}
	if (format == "jpeg" && filepath.Ext(source) != ".jpg") || (format == "png" && filepath.Ext(source) != ".png") {
		return errors.New("image format does not match its extension")
	}
	if config.Width <= 0 || config.Height <= 0 || config.Width > maxSourcePixels/config.Height {
		return fmt.Errorf("image dimensions %dx%d exceed the %d-pixel input limit or are invalid", config.Width, config.Height, maxSourcePixels)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	original, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("decode image: %w", err)
	}
	width, height := config.Width, config.Height
	var data []byte
	if width <= maxWidth && info.Size() <= maxBytes {
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return err
		}
		data, err = io.ReadAll(io.LimitReader(file, maxBytes+1))
	} else {
		data, width, height, err = resizeImage(original, format)
	}
	if err != nil {
		return err
	}
	if len(data) > maxBytes {
		return fmt.Errorf("result is %dx%d and %d bytes, over the %d-byte limit", width, height, len(data), maxBytes)
	}
	if err := os.WriteFile(destination, data, 0644); err != nil {
		return err
	}
	fmt.Fprintf(report, "%s: %dx%d, %d bytes -> %dx%d, %d bytes\n", filepath.Base(source), config.Width, config.Height, info.Size(), width, height, len(data))
	return nil
}

func resizeImage(original image.Image, format string) ([]byte, int, int, error) {
	bounds := original.Bounds()
	width := min(bounds.Dx(), maxWidth)
	var buffer bytes.Buffer
	for {
		height := max(1, int((int64(bounds.Dy())*int64(width)+int64(bounds.Dx()/2))/int64(bounds.Dx())))
		candidate := original
		if width != bounds.Dx() || height != bounds.Dy() {
			resized := image.NewNRGBA(image.Rect(0, 0, width, height))
			imageScaler(width, bounds.Dy()).Scale(resized, resized.Bounds(), original, bounds, draw.Src, nil)
			candidate = resized
		}
		buffer.Reset()
		var err error
		if format == "jpeg" {
			err = jpeg.Encode(&buffer, candidate, &jpeg.Options{Quality: 82})
		} else {
			encoder := png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(&buffer, candidate)
		}
		if err != nil {
			return nil, 0, 0, fmt.Errorf("encode image: %w", err)
		}
		// Below the preferred width, only the hard byte cap justifies another resize.
		if buffer.Len() <= targetBytes || (width <= preferredMinWidth && buffer.Len() <= maxBytes) || width == 1 {
			return buffer.Bytes(), width, height, nil
		}
		nextWidth := max(1, width*85/100)
		if width > preferredMinWidth {
			nextWidth = max(preferredMinWidth, nextWidth)
		}
		width = nextWidth
	}
}

func imageScaler(width, sourceHeight int) draw.Scaler {
	// The x/image kernel uses four float64 values per output column per source row.
	if int64(width)*int64(sourceHeight)*32 > maxKernelScratchBytes {
		return draw.ApproxBiLinear
	}
	return draw.CatmullRom
}
