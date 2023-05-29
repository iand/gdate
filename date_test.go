package gdate

import (
	"testing"
)

func TestSortsBefore(t *testing.T) {
	testCases := []struct {
		date      Date
		before    []Date
		notBefore []Date
	}{
		{
			date: &Precise{Y: 1845, M: 6, D: 15},
			before: []Date{
				&Precise{Y: 1845, M: 6, D: 16},
				&Precise{Y: 1846, M: 6, D: 15},
				&Precise{Y: 1845, M: 7, D: 15},
				&BeforeYear{Y: 1846},
				&AfterYear{Y: 1845},
				&AfterYear{Y: 1846},
				&AboutYear{Y: 1846},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1845, M: 7},
				&MonthYear{Y: 1846, M: 1},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&Precise{Y: 1845, M: 6, D: 14},
				&Precise{Y: 1845, M: 6, D: 15},
				&Precise{Y: 1844, M: 6, D: 15},
				&Precise{Y: 1845, M: 5, D: 15},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1845, M: 6},
				&MonthYear{Y: 1845, M: 5},
				&MonthYear{Y: 1844, M: 7},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &Precise{Y: 1845, M: 1, D: 1},
			notBefore: []Date{
				&YearQuarter{Y: 1845, Q: 1},
				&MonthYear{Y: 1845, M: 1},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &BeforeYear{Y: 1845},
			before: []Date{
				&BeforeYear{Y: 1846},
				&Precise{Y: 1845, M: 5, D: 15},
				&AfterYear{Y: 1845},
				&AfterYear{Y: 1846},
				&Year{Y: 1845},
				&AboutYear{Y: 1845},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&EstimatedYear{Y: 1845},
				&MonthYear{Y: 1845, M: 5},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&Precise{Y: 1844, M: 5, D: 15},
				&AfterYear{Y: 1844},
				&Year{Y: 1844},
				&AboutYear{Y: 1844},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1844, M: 12},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},

		{
			date: &AfterYear{Y: 1845},
			before: []Date{
				&AfterYear{Y: 1846},
				&Precise{Y: 1846, M: 5, D: 15},
				&BeforeYear{Y: 1846},
				&Year{Y: 1846},
				&AboutYear{Y: 1846},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1846, M: 5},
				&YearRange{Lower: 1846, Upper: 1847},
			},
			notBefore: []Date{
				&AfterYear{Y: 1845},
				&AfterYear{Y: 1844},
				&Precise{Y: 1845, M: 5, D: 15},
				&Precise{Y: 1844, M: 5, D: 15},
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&Year{Y: 1844},
				&Year{Y: 1845},
				&AboutYear{Y: 1845},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&EstimatedYear{Y: 1845},
				&MonthYear{Y: 1845, M: 5},
				&YearRange{Lower: 1845, Upper: 1847},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},

		{
			date: &Year{Y: 1845},
			before: []Date{
				&Year{Y: 1846},
				&Precise{Y: 1845, M: 6, D: 15},
				&Precise{Y: 1846, M: 6, D: 15},
				&BeforeYear{Y: 1846},
				&AfterYear{Y: 1845},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1846},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1846, M: 5},
				&MonthYear{Y: 1845, M: 5},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&Year{Y: 1844},
				&Year{Y: 1845},
				&Precise{Y: 1844, M: 6, D: 15},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&AboutYear{Y: 1844},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1844, M: 5},
				&YearQuarter{Y: 1844, Q: 4},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},

		{
			date: &AboutYear{Y: 1845},
			before: []Date{
				&AboutYear{Y: 1846},
				&Precise{Y: 1845, M: 6, D: 16},
				&Precise{Y: 1846, M: 6, D: 16},
				&BeforeYear{Y: 1846},
				&AfterYear{Y: 1845},
				&Year{Y: 1846},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1846, M: 1},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&Precise{Y: 1844, M: 6, D: 16},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1845, M: 12},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &YearQuarter{Y: 1845, Q: 1},
			before: []Date{
				&AboutYear{Y: 1846},
				&AfterYear{Y: 1845},
				&BeforeYear{Y: 1846},
				&Year{Y: 1846},
				&Precise{Y: 1845, M: 1, D: 1},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1845, M: 1},
				&MonthYear{Y: 1845, M: 4},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&Precise{Y: 1844, M: 12, D: 31},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &YearQuarter{Y: 1845, Q: 2},
			before: []Date{
				&AboutYear{Y: 1846},
				&AfterYear{Y: 1845},
				&BeforeYear{Y: 1846},
				&Year{Y: 1846},
				&Precise{Y: 1845, M: 4, D: 1},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1845, M: 4},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&Precise{Y: 1845, M: 1, D: 1},
				&Precise{Y: 1845, M: 3, D: 31},
				&Precise{Y: 1844, M: 12, D: 31},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1845, M: 3},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &YearQuarter{Y: 1845, Q: 3},
			before: []Date{
				&AboutYear{Y: 1846},
				&AfterYear{Y: 1845},
				&BeforeYear{Y: 1846},
				&Year{Y: 1846},
				&Precise{Y: 1845, M: 7, D: 1},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1845, M: 7},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&Precise{Y: 1845, M: 1, D: 1},
				&Precise{Y: 1845, M: 6, D: 30},
				&Precise{Y: 1844, M: 12, D: 31},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1845, M: 6},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &YearQuarter{Y: 1845, Q: 4},
			before: []Date{
				&AboutYear{Y: 1846},
				&AfterYear{Y: 1845},
				&BeforeYear{Y: 1846},
				&Year{Y: 1846},
				&Precise{Y: 1845, M: 10, D: 1},
				&EstimatedYear{Y: 1846},
				&MonthYear{Y: 1845, M: 10},
				&YearRange{Lower: 1850, Upper: 1860},
			},
			notBefore: []Date{
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&Precise{Y: 1845, M: 1, D: 1},
				&Precise{Y: 1845, M: 9, D: 30},
				&Precise{Y: 1844, M: 12, D: 31},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&MonthYear{Y: 1845, M: 9},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &EstimatedYear{Y: 1845},
			before: []Date{
				&EstimatedYear{Y: 1846},
				&AboutYear{Y: 1846},
				&Precise{Y: 1845, M: 6, D: 16},
				&Precise{Y: 1846, M: 6, D: 16},
				&BeforeYear{Y: 1846},
				&AfterYear{Y: 1845},
				&Year{Y: 1846},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&YearRange{Lower: 1850, Upper: 1850},
			},
			notBefore: []Date{
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&Precise{Y: 1844, M: 6, D: 16},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date:   &Unknown{},
			before: []Date{},
			notBefore: []Date{
				&EstimatedYear{Y: 1846},
				&AboutYear{Y: 1846},
				&Precise{Y: 1845, M: 6, D: 16},
				&Precise{Y: 1846, M: 6, D: 16},
				&BeforeYear{Y: 1846},
				&AfterYear{Y: 1845},
				&Year{Y: 1846},
				&YearQuarter{Y: 1845, Q: 1},
				&YearQuarter{Y: 1845, Q: 2},
				&YearQuarter{Y: 1845, Q: 3},
				&YearQuarter{Y: 1845, Q: 4},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&AboutYear{Y: 1845},
				&AboutYear{Y: 1844},
				&Precise{Y: 1844, M: 6, D: 16},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&Year{Y: 1845},
				&Year{Y: 1844},
				&YearRange{Lower: 1840, Upper: 1850},
			},
		},
		{
			date: &YearRange{Lower: 1840, Upper: 1850},
			before: []Date{
				&YearRange{Lower: 1841, Upper: 1850},
				&YearRange{Lower: 1840, Upper: 1849},
			},
			notBefore: []Date{
				&YearRange{Lower: 1840, Upper: 1851},
				&YearRange{Lower: 1839, Upper: 1850},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			for _, bd := range tc.before {
				before := SortsBefore(tc.date, bd)

				if !before {
					if ca, ok := tc.date.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", tc.date, ca.EarliestJulianDay(), ca.LatestJulianDay())
					}
					if cb, ok := bd.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", bd, cb.EarliestJulianDay(), cb.LatestJulianDay())
					}
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, bd, before, true)
				}

			}
			for _, nbd := range tc.notBefore {
				before := SortsBefore(tc.date, nbd)

				if before {
					if ca, ok := tc.date.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", tc.date, ca.EarliestJulianDay(), ca.LatestJulianDay())
					}
					if cb, ok := nbd.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", nbd, cb.EarliestJulianDay(), cb.LatestJulianDay())
					}
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, nbd, before, false)
				}

			}
		})
	}
}

func TestSortsBeforeWithStartOfYear(t *testing.T) {
	testCases := []struct {
		date      Date
		before    []Date
		notBefore []Date
	}{
		{
			date: &Precise{Y: 1760, M: 1, D: 1, C: Gregorian},
			before: []Date{
				&Precise{Y: 1760, M: 1, D: 2, C: Gregorian},
				&Precise{Y: 1760, M: 3, D: 26, C: Gregorian},
			},
			notBefore: []Date{
				&Precise{Y: 1759, M: 12, D: 31, C: Gregorian},
			},
		},
		{
			date: &Precise{Y: 1760, M: 1, D: 1, C: Julian},
			before: []Date{
				&Precise{Y: 1760, M: 1, D: 2, C: Julian},
				&Precise{Y: 1760, M: 3, D: 26, C: Julian},
			},
			notBefore: []Date{
				&Precise{Y: 1759, M: 12, D: 31, C: Julian},
			},
		},
		{
			date: &Precise{Y: 1750, M: 3, D: 24, C: Julian25Mar}, // last day of 1750
			before: []Date{
				&Precise{Y: 1751, M: 3, D: 25, C: Julian25Mar}, // first day of 1751
				&Precise{Y: 1751, M: 3, D: 24, C: Julian25Mar}, // last day of 1751
			},
			notBefore: []Date{
				&Precise{Y: 1750, M: 3, D: 25, C: Julian25Mar}, // first day of 1750
				&Precise{Y: 1750, M: 1, D: 1, C: Julian25Mar},  // towards the end of 1750
				&Precise{Y: 1749, M: 3, D: 24, C: Julian25Mar}, // last day of 1749
			},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			for _, bd := range tc.before {
				before := SortsBefore(tc.date, bd)

				if !before {
					if ca, ok := tc.date.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", tc.date, ca.EarliestJulianDay(), ca.LatestJulianDay())
					}
					if cb, ok := bd.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", bd, cb.EarliestJulianDay(), cb.LatestJulianDay())
					}
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, bd, before, true)
				}

			}
			for _, nbd := range tc.notBefore {
				before := SortsBefore(tc.date, nbd)

				if before {
					if ca, ok := tc.date.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", tc.date, ca.EarliestJulianDay(), ca.LatestJulianDay())
					}
					if cb, ok := nbd.(ComparableDate); ok {
						t.Logf("Julian Days for %q, earliest=%d, latest=%d", nbd, cb.EarliestJulianDay(), cb.LatestJulianDay())
					}
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, nbd, before, false)
				}

			}
		})
	}
}
