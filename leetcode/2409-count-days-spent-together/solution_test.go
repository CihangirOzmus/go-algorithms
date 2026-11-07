package countdaysspenttogether

import "testing"

func TestCountDaysTogether(t *testing.T) {
	tests := []struct {
		name                                         string
		arriveAlice, leaveAlice, arriveBob, leaveBob string
		want                                         int
	}{
		{"example 1", "08-15", "08-18", "08-16", "08-19", 3},
		{"example 2", "10-01", "10-31", "11-01", "12-31", 0},
		{"across months", "01-30", "03-02", "02-27", "12-31", 4},
		{"same single day", "05-05", "05-05", "05-05", "05-05", 1},
		{"one contains other", "01-01", "12-31", "06-01", "06-30", 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDaysTogether(tt.arriveAlice, tt.leaveAlice, tt.arriveBob, tt.leaveBob); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
