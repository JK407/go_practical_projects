package utils

import "testing"

func TestParseSize(t *testing.T) {
	sizeList := []string{"100B", "100KB", "300MB", "300GB", "100TB", "100PB"}
	for _, size := range sizeList {
		byteNum, sizeStr := ParseSize(size)
		t.Logf("byteNum: %d, sizeStr: %s", byteNum, sizeStr)
	}
}
