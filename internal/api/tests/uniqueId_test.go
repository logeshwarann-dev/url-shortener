package tests

import (
	"testing"

	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

func TestEncodeUrl(t *testing.T) {
	shortCode := utils.GetShortCode(6)
	t.Logf("Generated ShortCode: %v", shortCode)
	if len(shortCode) < 6 || len(shortCode) > 10 {
		t.Fatalf("Invalid Short Code: %v", shortCode)
	}
	t.Log("Short code is valid!")
}

func TestDecodeUrl(t *testing.T) {
	shortCode := "7ZSnOAFlLZeqiXuUI2VveNnLEna8heP1rKOnuCI64zb2D8UGropTz30OHFl46Eppi675umeUCgyMinKbWDyP5u7Je6m3z"
	longUrl, err := utils.DecodeString(shortCode)
	if err != nil {
		t.Errorf("error while decoding short code: %v", err.Error())
		t.Fatal("Decoding failed!")
	}
	t.Logf("Decoded URL: %v", longUrl)
	if len(longUrl) == 0 || len(longUrl) < 5 {
		t.Fatalf("Invalid decoded URL: %v", longUrl)
	}
	t.Log("long Url is valid!")

}
