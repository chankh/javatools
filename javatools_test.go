package javatools

import (
	"testing"
)

func TestParseJavaVersion(t *testing.T) {
	cases := []struct {
		in   string
		want JavaVersion
	}{
		{"1.8.0_66-b18", JavaVersion{1, 8, 0, 66, 18}},
		{"1.7.0_67-b25", JavaVersion{1, 7, 0, 67, 25}},
	}
	for _, c := range cases {
		got, err := ParseJavaVersion(c.in)
		if err != nil {
			t.Errorf("ParseJavaVersion(%q) returns error %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("ParseJavaVersion(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
