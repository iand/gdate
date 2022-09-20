package gdate_test

import (
	"fmt"
	"sort"

	"github.com/iand/gdate"
)

func Example() {
	input := []string{
		"7 Nov 1880",
		"5 Nov 1878",
		"6 Apr 1877",
		"before 1877",
		"not a date",
		"after 1878",
		"1885",
		"about 1879",
	}

	dates := []gdate.Date{}

	for _, in := range input {
		dt, _ := gdate.Parse(in)
		dates = append(dates, dt)
	}

	sort.Slice(dates, func(i, j int) bool {
		return gdate.SortsBefore(dates[i], dates[j])
	})

	fmt.Println("sorted dates")
	for _, dt := range dates {
		fmt.Printf(" - %s\n", dt.String())
	}

	// Output:
	// sorted dates
	//  - bef. 1877
	//  - 6 Apr 1877
	//  - 5 Nov 1878
	//  - aft. 1878
	//  - abt. 1879
	//  - 7 Nov 1880
	//  - 1885
	//  - not a date
}
