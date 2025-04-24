package gdate

import (
	"fmt"
	"testing"
)

func TestCalendarFmtYear(t *testing.T) {
	testCases := []struct {
		c       Calendar
		y, m, d int
		want    string
	}{
		{
			c:    Gregorian,
			y:    1400,
			m:    2,
			d:    5,
			want: "1400",
		},
		{
			c:    Gregorian,
			y:    1700,
			m:    12,
			d:    15,
			want: "1700",
		},
		{
			c:    Gregorian,
			y:    1800,
			m:    12,
			d:    15,
			want: "1800",
		},
		{
			c:    Gregorian,
			y:    800,
			m:    12,
			d:    15,
			want: "800",
		},
		{
			c:    Julian,
			y:    1,
			m:    1,
			d:    1,
			want: "1",
		},
		{
			c:    Julian,
			y:    1752,
			m:    1,
			d:    1,
			want: "1752",
		},
		{
			c:    Julian,
			y:    1752,
			m:    3,
			d:    35,
			want: "1752",
		},
		{
			c:    Julian25Mar,
			y:    1,
			m:    1,
			d:    1,
			want: "1/2",
		},
		{
			c:    Julian25Mar,
			y:    1752,
			m:    1,
			d:    1,
			want: "1752/3",
		},
		{
			c:    Julian25Mar,
			y:    1752,
			m:    3,
			d:    25,
			want: "1752",
		},
		{
			c:    Julian25Mar,
			y:    1800,
			m:    12,
			d:    1,
			want: "1800",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%02d_%02d_%04d", tc.c, tc.d, tc.m, tc.y), func(t *testing.T) {
			got := tc.c.FmtYear(tc.y, tc.m, tc.d)
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCalendarJulianDay(t *testing.T) {
	testCases := []struct {
		c       Calendar
		y, m, d int
		want    int
	}{
		{
			c:    Julian,
			y:    125,
			m:    4,
			d:    24,
			want: 1766828,
		},
		{
			c:    Julian,
			y:    1400,
			m:    7,
			d:    4,
			want: 2232593,
		},
		{
			c:    Julian,
			y:    1752,
			m:    3,
			d:    24,
			want: 2361059,
		},
		{
			c:    Julian,
			y:    1752,
			m:    3,
			d:    25,
			want: 2361060,
		},
		{
			c:    Julian,
			y:    1900,
			m:    3,
			d:    24,
			want: 2415116,
		},
		{
			c:    Gregorian,
			y:    1900,
			m:    3,
			d:    24,
			want: 2415102,
		},
		{
			c:    Gregorian,
			y:    1900,
			m:    3,
			d:    25,
			want: 2415103,
		},
		{
			c:    Gregorian,
			y:    1582,
			m:    10,
			d:    15,
			want: 2299160,
		},
		{
			c:    Julian25Mar,
			y:    1400,
			m:    7,
			d:    4,
			want: 2232593,
		},
		{
			c:    Julian25Mar,
			y:    1752,
			m:    3,
			d:    24,
			want: 2361059,
		},
		{
			c:    Julian25Mar,
			y:    1752,
			m:    3,
			d:    25,
			want: 2361060,
		},
		{
			c:    Julian25Mar,
			y:    1900,
			m:    3,
			d:    24,
			want: 2415116,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%02d_%02d_%04d", tc.c, tc.d, tc.m, tc.y), func(t *testing.T) {
			got := tc.c.JulianDay(tc.y, tc.m, tc.d)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
