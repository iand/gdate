package gdate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		s    string
		alts []string
		err  bool
		want Date
	}{
		{
			s:    "2 Apr 1871",
			alts: []string{"2 Apr, 1871", "2 April 1871", "2 April, 1871", "02 April 1871", "Apr 2, 1871", "Apr 2 1871", "Apr 02 1871"},
			want: &Precise{Y: 1871, M: 4, D: 2},
		},
		{
			s:    "23 Apr 1871",
			want: &Precise{Y: 1871, M: 4, D: 23},
		},
		{
			s:    "bef 1950",
			alts: []string{"bef. 1950", "before 1950", "BEF 1950", "BEF. 1950", "BEFORE 1950"},
			want: &BeforeYear{Y: 1950},
		},
		{
			s:    "aft 1950",
			alts: []string{"aft. 1950", "after 1950", "AFT 1950", "AFT. 1950", "AFTER 1950"},
			want: &AfterYear{Y: 1950},
		},
		{
			s:    "",
			want: &Unknown{},
		},
		{
			s:    "some time in 1920",
			want: &Unknown{Text: "some time in 1920"},
		},
		{
			s:    "1950",
			want: &Year{Y: 1950},
		},
		{
			s:    "about 1950",
			alts: []string{"abt. 1950"},
			want: &AboutYear{Y: 1950},
		},
		{
			s:    "January 1950",
			alts: []string{"january 1950"},
			want: &MonthYear{Y: 1950, M: 1},
		},
		{
			s:    "February 1950",
			alts: []string{"february 1950"},
			want: &MonthYear{Y: 1950, M: 2},
		},
		{
			s:    "March 1950",
			alts: []string{"march 1950"},
			want: &MonthYear{Y: 1950, M: 3},
		},
		{
			s:    "April 1950",
			alts: []string{"april 1950"},
			want: &MonthYear{Y: 1950, M: 4},
		},
		{
			s:    "May 1950",
			alts: []string{"may 1950"},
			want: &MonthYear{Y: 1950, M: 5},
		},
		{
			s:    "June 1950",
			alts: []string{"june 1950"},
			want: &MonthYear{Y: 1950, M: 6},
		},
		{
			s:    "July 1950",
			alts: []string{"july 1950"},
			want: &MonthYear{Y: 1950, M: 7},
		},
		{
			s:    "August 1950",
			alts: []string{"august 1950"},
			want: &MonthYear{Y: 1950, M: 8},
		},
		{
			s:    "September 1950",
			alts: []string{"september 1950"},
			want: &MonthYear{Y: 1950, M: 9},
		},
		{
			s:    "October 1950",
			alts: []string{"october 1950"},
			want: &MonthYear{Y: 1950, M: 10},
		},
		{
			s:    "November 1950",
			alts: []string{"november 1950"},
			want: &MonthYear{Y: 1950, M: 11},
		},
		{
			s:    "December 1950",
			alts: []string{"december 1950"},
			want: &MonthYear{Y: 1950, M: 12},
		},
		{
			s:    "1920-1923",
			want: &YearRange{Lower: 1920, Upper: 1923},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			dt, err := Parse(tc.s)
			if err != nil && !tc.err {
				t.Fatalf("got unexpected error: %v", err)
			}
			if err == nil && tc.err {
				t.Fatalf("missing expected error")
			}

			if diff := cmp.Diff(tc.want, dt); diff != "" {
				t.Errorf("Parse(%q) mismatch (-want +got):\n%s", tc.s, diff)
			}

			for _, alt := range tc.alts {
				dt, err := Parse(alt)
				if err != nil && !tc.err {
					t.Fatalf("got unexpected error: %v", err)
				}
				if err == nil && tc.err {
					t.Fatalf("missing expected error")
				}

				if diff := cmp.Diff(tc.want, dt); diff != "" {
					t.Errorf("Parse(%q) mismatch (-want +got):\n%s", alt, diff)
				}
			}
		})
	}
}

func TestParseAssumeGROQuarter(t *testing.T) {
	testCases := []struct {
		s    string
		alts []string
		err  bool
		want Date
	}{
		{
			s:    "Mar 1950",
			alts: []string{"mar 1950", "MAR 1950", "Q1 1950", "1950Q1", "jan 1950", "JAN 1950"},
			want: &YearQuarter{Y: 1950, Q: 1},
		},
		{
			s:    "Jun 1950",
			alts: []string{"jun 1950", "JUN 1950", "Q2 1950", "1950Q2", "apr 1950", "APR 1950"},
			want: &YearQuarter{Y: 1950, Q: 2},
		},
		{
			s:    "Sep 1950",
			alts: []string{"sep 1950", "SEP 1950", "Q3 1950", "1950Q3", "jul 1950", "JUL 1950"},
			want: &YearQuarter{Y: 1950, Q: 3},
		},
		{
			s:    "Dec 1950",
			alts: []string{"dec 1950", "DEC 1950", "Q4 1950", "1950Q4", "oct 1950", "OCT 1950"},
			want: &YearQuarter{Y: 1950, Q: 4},
		},
		{
			s:    "February 1950",
			alts: []string{"february 1950"},
			want: &MonthYear{Y: 1950, M: 2},
		},
		{
			s:    "May 1950",
			alts: []string{"may 1950"},
			want: &MonthYear{Y: 1950, M: 5},
		},
		{
			s:    "August 1950",
			alts: []string{"august 1950"},
			want: &MonthYear{Y: 1950, M: 8},
		},
		{
			s:    "November 1950",
			alts: []string{"november 1950"},
			want: &MonthYear{Y: 1950, M: 11},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			p := &Parser{
				AssumeGROQuarter: true,
			}

			dt, err := p.Parse(tc.s)
			if err != nil && !tc.err {
				t.Fatalf("got unexpected error: %v", err)
			}
			if err == nil && tc.err {
				t.Fatalf("missing expected error")
			}

			if diff := cmp.Diff(tc.want, dt); diff != "" {
				t.Errorf("Parse(%q) mismatch (-want +got):\n%s", tc.s, diff)
			}

			for _, alt := range tc.alts {
				dt, err := p.Parse(alt)
				if err != nil && !tc.err {
					t.Fatalf("got unexpected error: %v", err)
				}
				if err == nil && tc.err {
					t.Fatalf("missing expected error")
				}

				if diff := cmp.Diff(tc.want, dt); diff != "" {
					t.Errorf("Parse(%q) mismatch (-want +got):\n%s", alt, diff)
				}
			}
		})
	}
}

func TestParseReckoningLocation(t *testing.T) {
	testCases := []struct {
		s    string
		l    ReckoningLocation
		want Date
	}{
		{
			s:    "2 Apr 1752",
			l:    EnglandAndWales,
			want: &Precise{Y: 1752, M: 4, D: 2, C: Gregorian},
		},
		{
			s:    "2 Apr 1751",
			l:    EnglandAndWales,
			want: &Precise{Y: 1751, M: 4, D: 2, C: Julian25Mar},
		},
		{
			s:    "2 Apr 1752",
			l:    Scotland,
			want: &Precise{Y: 1752, M: 4, D: 2, C: Gregorian},
		},
		{
			s:    "2 Apr 1751",
			l:    Scotland,
			want: &Precise{Y: 1751, M: 4, D: 2, C: Julian},
		},
		{
			s:    "2 Apr 1558",
			l:    Scotland,
			want: &Precise{Y: 1558, M: 4, D: 2, C: Julian25Mar},
		},
		{
			s:    "2 Apr 1752",
			l:    Ireland,
			want: &Precise{Y: 1752, M: 4, D: 2, C: Gregorian},
		},
		{
			s:    "2 Apr 1751",
			l:    Ireland,
			want: &Precise{Y: 1751, M: 4, D: 2, C: Julian25Mar},
		},
		{
			s:    "bef 1752",
			l:    EnglandAndWales,
			want: &BeforeYear{Y: 1752, C: Julian25Mar},
		},
		{
			s:    "bef 1753",
			l:    EnglandAndWales,
			want: &BeforeYear{Y: 1753, C: Gregorian},
		},
		{
			s:    "aft 1751",
			l:    EnglandAndWales,
			want: &AfterYear{Y: 1751, C: Gregorian},
		},
		{
			s:    "aft 1750",
			l:    EnglandAndWales,
			want: &AfterYear{Y: 1750, C: Julian25Mar},
		},
		{
			s:    "",
			l:    EnglandAndWales,
			want: &Unknown{C: Gregorian},
		},
		{
			s:    "1751",
			l:    EnglandAndWales,
			want: &Year{Y: 1751, C: Julian25Mar},
		},
		{
			s:    "1752",
			l:    EnglandAndWales,
			want: &Year{Y: 1752, C: Gregorian},
		},
		{
			s:    "about 1751",
			l:    EnglandAndWales,
			want: &AboutYear{Y: 1751, C: Julian25Mar},
		},
		{
			s:    "about 1752",
			l:    EnglandAndWales,
			want: &AboutYear{Y: 1752, C: Gregorian},
		},
		{
			s:    "January 1751",
			l:    EnglandAndWales,
			want: &MonthYear{Y: 1751, M: 1, C: Julian25Mar},
		},
		{
			s:    "January 1752",
			l:    EnglandAndWales,
			want: &MonthYear{Y: 1752, M: 1, C: Gregorian},
		},
		{
			s:    "1751-1753",
			want: &YearRange{Lower: 1751, Upper: 1753, C: Julian25Mar},
		},
		{
			s:    "1752-1753",
			want: &YearRange{Lower: 1752, Upper: 1753, C: Gregorian},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			p := &Parser{
				ReckoningLocation: tc.l,
			}
			dt, err := p.Parse(tc.s)
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.want, dt); diff != "" {
				t.Errorf("Parse(%q) mismatch (-want +got):\n%s", tc.s, diff)
			}
		})
	}
}
