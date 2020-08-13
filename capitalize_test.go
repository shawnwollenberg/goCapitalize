package capitalize

import "testing"

func TestCapitalize(t *testing.T) {
	want := "Shawn Is Cool"
	if got := capitalize("shawn is cool"); got != want {
		t.Errorf("capitalize() = %q, want %q", got, want)
	}
}
