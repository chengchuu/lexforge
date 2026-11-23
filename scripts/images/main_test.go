package main

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/image/draw"
)

func TestOptimizeDirectory(t *testing.T) {
	input := t.TempDir()
	output := filepath.Join(t.TempDir(), "output")
	picture := image.NewNRGBA(image.Rect(0, 0, 640, 480))
	picture.SetNRGBA(320, 240, color.NRGBA{R: 255, A: 255})
	var large, wideJPEG, small bytes.Buffer
	if err := png.Encode(&large, picture); err != nil {
		t.Fatal(err)
	}
	if err := jpeg.Encode(&small, image.NewRGBA(image.Rect(0, 0, 160, 240)), nil); err != nil {
		t.Fatal(err)
	}
	if err := jpeg.Encode(&wideJPEG, picture, nil); err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string][]byte{"large.png": large.Bytes(), "small.jpg": small.Bytes(), "wide.jpg": wideJPEG.Bytes()} {
		if err := os.WriteFile(filepath.Join(input, name), data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := optimizeDirectory(input, output, io.Discard); err != nil {
		t.Fatal(err)
	}
	resized, err := os.ReadFile(filepath.Join(output, "large.png"))
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := png.Decode(bytes.NewReader(resized))
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Bounds() != image.Rect(0, 0, 200, 150) || len(resized) > maxBytes {
		t.Fatalf("unexpected output: %v, %d bytes", decoded.Bounds(), len(resized))
	}
	if _, _, _, alpha := decoded.At(0, 0).RGBA(); alpha != 0 {
		t.Fatal("resized PNG lost transparency")
	}
	jpegOutput, err := os.ReadFile(filepath.Join(output, "wide.jpg"))
	if err != nil {
		t.Fatal(err)
	}
	jpegConfig, err := jpeg.DecodeConfig(bytes.NewReader(jpegOutput))
	if err != nil || jpegConfig.Width != 200 || jpegConfig.Height != 150 || len(jpegOutput) > maxBytes {
		t.Fatalf("unexpected JPEG output: %+v, %d bytes, %v", jpegConfig, len(jpegOutput), err)
	}
	copied, err := os.ReadFile(filepath.Join(output, "small.jpg"))
	if err != nil || !bytes.Equal(copied, small.Bytes()) {
		t.Fatal("already compliant portrait JPEG changed", err)
	}
	original, err := os.ReadFile(filepath.Join(input, "large.png"))
	if err != nil || !bytes.Equal(original, large.Bytes()) {
		t.Fatal("input changed", err)
	}
	if err := optimizeDirectory(input, output, io.Discard); !errors.Is(err, os.ErrExist) {
		t.Fatalf("expected existing directory error, got %v", err)
	}
}

func TestOptimizeDirectoryRejectsMismatchedFormat(t *testing.T) {
	input := t.TempDir()
	output := filepath.Join(t.TempDir(), "output")
	var picture bytes.Buffer
	if err := png.Encode(&picture, image.NewNRGBA(image.Rect(0, 0, 1, 1))); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"a.png", "b.jpg"} {
		if err := os.WriteFile(filepath.Join(input, name), picture.Bytes(), 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := optimizeDirectory(input, output, io.Discard); err == nil || !strings.Contains(err.Error(), "b.jpg: image format does not match") {
		t.Fatalf("expected mismatch error, got %v", err)
	}
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		t.Fatalf("invalid batch left an output directory: %v", err)
	}
}

func TestResizeBelowPreferredWidthToMeetByteCap(t *testing.T) {
	input := t.TempDir()
	output := filepath.Join(t.TempDir(), "output")
	picture := image.NewNRGBA(image.Rect(0, 0, 100, 1200))
	rng := rand.New(rand.NewSource(7))
	for i := 0; i < len(picture.Pix); i += 4 {
		rgb := rng.Uint32()
		picture.Pix[i] = byte(rgb)
		picture.Pix[i+1] = byte(rgb >> 8)
		picture.Pix[i+2] = byte(rgb >> 16)
		picture.Pix[i+3] = 255
	}
	var source bytes.Buffer
	if err := png.Encode(&source, picture); err != nil {
		t.Fatal(err)
	}
	if source.Len() <= maxBytes {
		t.Fatalf("fixture must exceed %d bytes", maxBytes)
	}
	if err := os.WriteFile(filepath.Join(input, "tall.png"), source.Bytes(), 0600); err != nil {
		t.Fatal(err)
	}
	if err := optimizeDirectory(input, output, io.Discard); err != nil {
		t.Fatal(err)
	}
	result, err := os.ReadFile(filepath.Join(output, "tall.png"))
	if err != nil {
		t.Fatal(err)
	}
	config, err := png.DecodeConfig(bytes.NewReader(result))
	if err != nil || config.Width >= 100 || config.Height != config.Width*12 || len(result) > maxBytes {
		t.Fatalf("unexpected output: %+v, %d bytes, %v", config, len(result), err)
	}
}

func TestTallImagesUseBoundedScratchScaler(t *testing.T) {
	if imageScaler(200, 125_000) != draw.ApproxBiLinear {
		t.Fatal("tall image should use the scaler without the kernel scratch buffer")
	}
	if imageScaler(200, 480) != draw.CatmullRom {
		t.Fatal("ordinary image should retain Catmull-Rom scaling")
	}
	tall := image.NewNRGBA(image.Rect(0, 0, 201, 3500))
	_, width, height, err := resizeImage(tall, "png")
	if err != nil || width != 200 || height != 3483 {
		t.Fatalf("unexpected tall output: %dx%d, %v", width, height, err)
	}
}
