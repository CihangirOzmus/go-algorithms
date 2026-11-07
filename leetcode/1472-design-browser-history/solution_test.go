package designbrowserhistory

import "testing"

func TestBrowserHistoryExample(t *testing.T) {
	b := Constructor("leetcode.com")
	b.Visit("google.com")
	b.Visit("facebook.com")
	b.Visit("youtube.com")

	steps := []struct {
		op   string
		n    int
		want string
	}{
		{"back", 1, "facebook.com"},
		{"back", 1, "google.com"},
		{"forward", 1, "facebook.com"},
		{"visit", 0, "linkedin.com"}, // clears forward history
		{"forward", 2, "linkedin.com"},
		{"back", 2, "google.com"},
		{"back", 7, "leetcode.com"},
	}
	for i, s := range steps {
		var got string
		switch s.op {
		case "back":
			got = b.Back(s.n)
		case "forward":
			got = b.Forward(s.n)
		case "visit":
			b.Visit(s.want)
			continue
		}
		if got != s.want {
			t.Fatalf("step %d: %s(%d) = %q, want %q", i, s.op, s.n, got, s.want)
		}
	}
}
