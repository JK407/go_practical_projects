package utils

import "testing"

func TestGetValueSize(t *testing.T) {
	m := []interface{}{
		"oberl", // 16+5
		18,
		[]int{1, 2, 3},
	}
	for _, v := range m {
		t.Logf("value: %#v, size: %d", v, GetValueSize(v))
	}
}
