package utils

import "testing"

func TestReverse_ForUnicodeString(t *testing.T) {
	expected := "✌️ఘஹআՖØĈ🙏"
	actual := Reverse("🙏ĈØՖআஹఘ✌️")
	if expected != actual {
		t.Errorf("Reverse() = %v, want %v", actual, expected)
	}
}