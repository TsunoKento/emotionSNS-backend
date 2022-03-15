package pkg

import "testing"

func TestRandString(t *testing.T) {
	n := 12

	r := RandString(12)

	if len(r) != n {
		t.Fatal("指定された数値以外の数字が返却されています")
	}
}

func TestRandString02(t *testing.T) {
	t.Fatal("Error")
}
