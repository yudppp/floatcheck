package floatcheck

import "testing"

func TestContainsFloatSpecifier(t *testing.T) {
	cases := []struct {
		format string
		want   bool
	}{
		{"%f", true},
		{"%F", true},
		{"%g", true},
		{"%G", true},
		{"%e", true},
		{"%E", true},
		{"%d", false},
		{"%s", false},
		{"abc", false},
	}
	for _, c := range cases {
		if got := containsFloatSpecifier(c.format); got != c.want {
			t.Errorf("containsFloatSpecifier(%q) = %v, want %v", c.format, got, c.want)
		}
	}
}
