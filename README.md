# gdate

A Go package to parse and order genealogical dates.

[![Check Status](https://github.com/iand/gdate/actions/workflows/check.yml/badge.svg)](https://github.com/iand/gdate/actions/workflows/check.yml)
[![Test Status](https://github.com/iand/gdate/actions/workflows/test.yml/badge.svg)](https://github.com/iand/gdate/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iand/gdate)](https://goreportcard.com/report/github.com/iand/gdate)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/gdate)



## Overview

The following types of date are supported:

 - Precise, dates that have a known year, month and day
 - Year, dates that occur within a known specific year
 - BeforeYear, dates that are before the start of a specific year, often written as "bef. 1905"
 - AfterYear, dates that are after the start of a specific year, often written as "aft. 1921"
 - AboutYear, dates that are near to a specific year, often written as "abt. 1850"
 - YearQuarter, a year plus the quarter according to the UK general register office convention, 
 - EstimatedYear, dates that are estimated to be a specific year, often written as "est. 1960"
 - Unknown, an unknown date

## Usage

An example of usage:

```Golang
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


```





## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.

