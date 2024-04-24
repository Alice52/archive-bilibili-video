package api

import (
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/micro-services-roadmap/kit-common/viperx/initialize"
	"github.com/skip2/go-qrcode"
	"os"
	"testing"
)

func TestQrCodeImage(t *testing.T) {
	defer func(fn string) {
		os.Remove(fn)
	}("qr.png")

	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		return
	}
}

func TestGenerateAndEmail(t *testing.T) {
	initialize.Viper(&global.CONFIG, "../config-local.yaml")

	err := GenerateAndEmail("zack", qrcode.Low, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
