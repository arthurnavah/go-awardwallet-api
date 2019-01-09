package awardwallet

import (
	"testing"
)

func TestClient(t *testing.T) {
	award := NewClient("https://business.awardwallet.com/api/export/v1/account/1", "1010")

	if award == nil {
		t.Error()
	}
}
