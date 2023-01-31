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
