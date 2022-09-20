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
			},
			notBefore: []Date{
				&BeforeYear{Y: 1845},
				&BeforeYear{Y: 1844},
				&Precise{Y: 1844, M: 5, D: 15},
				&AfterYear{Y: 1844},
				&Year{Y: 1844},
				&AboutYear{Y: 1844},
				&EstimatedYear{Y: 1844},
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
				&AboutYear{Y: 1846},
				&EstimatedYear{Y: 1846},
			},
			notBefore: []Date{
				&Year{Y: 1844},
				&Year{Y: 1845},
				&Precise{Y: 1844, M: 6, D: 15},
				&BeforeYear{Y: 1845},
				&AfterYear{Y: 1844},
				&AboutYear{Y: 1844},
				&AboutYear{Y: 1845},
				&EstimatedYear{Y: 1844},
				&EstimatedYear{Y: 1845},
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
				&Precise{Y: 1845, M: 9, D: 31},
				&Precise{Y: 1844, M: 12, D: 31},
				&EstimatedYear{Y: 1845},
				&EstimatedYear{Y: 1844},
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
			},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			for _, bd := range tc.before {
				before := SortsBefore(tc.date, bd)

				if !before {
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, bd, before, true)
				}

			}
			for _, nbd := range tc.notBefore {
				before := SortsBefore(tc.date, nbd)

				if before {
					t.Errorf("got SortsBefore(%q,%q)=%v, wanted %v", tc.date, nbd, before, false)
				}

			}
		})
	}
}
