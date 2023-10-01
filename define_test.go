package basebodyig

import "testing"

func TestCodingRange(t *testing.T) {
	前加字数 := len([]rune(BodYig前加字))
	if 前加字数 != 8 {
		t.Fatal("excpeted 前加字数 ==", 8, "but got", 前加字数)
	}
	后加字数 := len([]rune(BodYig后加字))
	if 后加字数 != 16 {
		t.Fatal("excpeted 后加字数 ==", 16, "but got", 后加字数)
	}
	再后加字数 := len([]rune(BodYig再后加字))
	if 再后加字数 != 4 {
		t.Fatal("excpeted 再后加字数 ==", 4, "but got", 再后加字数)
	}
}
