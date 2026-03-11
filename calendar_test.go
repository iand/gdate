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
		// Dual-date suffix: always two digits of the NS year (OS year y, NS year y+1).
		{
			c:    Julian25Mar,
			y:    1650,
			m:    3,
			d:    10,
			want: "1650/51",
		},
		{
			c:    Julian25Mar,
			y:    1649,
			m:    1,
			d:    15,
			want: "1649/50",
		},
		{
			c:    Julian25Mar,
			y:    1752,
			m:    1,
			d:    1,
			want: "1752/53",
		},
		{
			c:    Julian25Mar,
			y:    1699,
			m:    2,
			d:    1,
			want: "1699/00",
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
			want: 2415103,
		},
		{
			c:    Gregorian,
			y:    1900,
			m:    3,
			d:    25,
			want: 2415104,
		},
		{
			c:    Gregorian,
			y:    1582,
			m:    10,
			d:    15,
			want: 2299161, // standard reference: first day of Gregorian calendar
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
			d:    25,
			want: 2361060,
		},
		// OS year 1650, June 4 (mid-year, no year adjustment): same as Julian(1650, 6, 4)
		{
			c:    Julian25Mar,
			y:    1650,
			m:    6,
			d:    4,
			want: 2323875,
		},
		// OS year 1650, March 10 (early March, before 25th): maps to Julian(1651, 3, 10).
		// Must sort AFTER June 4, 1650.
		{
			c:    Julian25Mar,
			y:    1650,
			m:    3,
			d:    10,
			want: 2324154,
		},
		// OS year 1752, March 24 (before 25th): maps to Julian(1753, 3, 24).
		{
			c:    Julian25Mar,
			y:    1752,
			m:    3,
			d:    24,
			want: 2361424,
		},
		{
			c:    Julian25Mar,
			y:    1900,
			m:    3,
			d:    24,
			want: 2415481,
		},
	}

	// Verify sort-order invariant: Julian25Mar(1650, 3, 10) must sort after
	// Julian25Mar(1650, 6, 4). March 10 OS-1650 = Julian March 10, 1651, which
	// is chronologically later than June 4, 1650.
	t.Run("sort_order_early_march_after_june", func(t *testing.T) {
		march10 := Julian25Mar.JulianDay(1650, 3, 10)
		june4 := Julian25Mar.JulianDay(1650, 6, 4)
		if march10 <= june4 {
			t.Errorf("Julian25Mar(1650,3,10)=%d should be > Julian25Mar(1650,6,4)=%d", march10, june4)
		}
	})

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%02d_%02d_%04d", tc.c, tc.d, tc.m, tc.y), func(t *testing.T) {
			got := tc.c.JulianDay(tc.y, tc.m, tc.d)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
