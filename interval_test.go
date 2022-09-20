package gdate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntervalBetween(t *testing.T) {
	testCases := []struct {
		a       Date
		b       Date
		unknown bool
		want    Interval
	}{
		{
			a:    &Unknown{},
			b:    &Unknown{},
			want: &UnknownInterval{},
		},
		{
			a:    &Precise{Y: 1845, M: 6, D: 16},
			b:    &Unknown{},
			want: &UnknownInterval{},
		},
		{
			a:    &Unknown{},
			b:    &Precise{Y: 1845, M: 6, D: 16},
			want: &UnknownInterval{},
		},
		{
			a:    &Precise{Y: 1845, M: 6, D: 15},
			b:    &Precise{Y: 1845, M: 6, D: 16},
			want: &PreciseInterval{D: 1},
		},
		{
			a:    &Precise{Y: 1845, M: 6, D: 15},
			b:    &Precise{Y: 1845, M: 7, D: 16},
			want: &PreciseInterval{M: 1, D: 1},
		},
		{
			a:    &Precise{Y: 1845, M: 6, D: 15},
			b:    &Precise{Y: 1845, M: 7, D: 15},
			want: &PreciseInterval{M: 1, D: 0},
		},
		{
			a:    &Precise{Y: 1845, M: 6, D: 15},
			b:    &Precise{Y: 1845, M: 7, D: 14},
			want: &PreciseInterval{M: 0, D: 29},
		},
		{
			a:    &Precise{Y: 1845, M: 5, D: 15},
			b:    &Precise{Y: 1845, M: 6, D: 14},
			want: &PreciseInterval{M: 0, D: 30},
		},
		{
			// non leap year
			a:    &Precise{Y: 1905, M: 2, D: 27},
			b:    &Precise{Y: 1905, M: 3, D: 1},
			want: &PreciseInterval{M: 0, D: 2},
		},
		{
			// leap year
			a:    &Precise{Y: 1904, M: 2, D: 27},
			b:    &Precise{Y: 1904, M: 3, D: 1},
			want: &PreciseInterval{M: 0, D: 3},
		},

		// Years
		{
			a:    &Year{Y: 1845},
			b:    &Year{Y: 1845},
			want: &YearsInterval{Y: 0},
		},
		{
			a:    &Year{Y: 1840},
			b:    &Year{Y: 1845},
			want: &YearsInterval{Y: 5},
		},
		{
			a:    &Year{Y: 1846},
			b:    &Year{Y: 1845},
			want: &YearsInterval{Y: 1},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			in := IntervalBetween(tc.a, tc.b)

			if diff := cmp.Diff(tc.want, in); diff != "" {
				t.Errorf("IntervalBetween mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
