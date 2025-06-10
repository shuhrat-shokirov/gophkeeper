package utils

import "testing"

func TestGenerateShortUUID(t *testing.T) {
	uuid := GenerateShortUUID()
	if len(uuid) == 0 {
		t.Errorf("GenerateShortUUID() returned an empty string")
		return
	}
}
