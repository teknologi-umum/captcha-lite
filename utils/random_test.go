package utils_test

import (
	"captcha-lite/utils"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	n := utils.GenerateRandomNumber()
	if len(n) != 3 {
		t.Errorf("GenerateRandomNumber() should return 3 digits, got %d", len(n))
	}
}
