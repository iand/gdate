package gdate

import (
	"regexp"
	"strconv"
	"time"
)

var (
	reYear        = regexp.MustCompile(`^\d{4}$`)
	reBeforeYear  = regexp.MustCompile(`(?i)^bef(?:.|ore)?\s+(\d{4})\s*$`)
	reAfterYear   = regexp.MustCompile(`(?i)^aft(?:.|er)?\s+(\d{4})\s*$`)
	reAboutYear   = regexp.MustCompile(`(?i)^(?:abt|abt.|about)?\s+(\d{4})\s*$`)
	reMarQuarter  = regexp.MustCompile(`(?i)^(?:mar|mar.|march|q1|jan)?\s+(\d{4})\s*$`)
	reJunQuarter  = regexp.MustCompile(`(?i)^(?:jun|jun.|june|q2|apr)?\s+(\d{4})\s*$`)
	reSepQuarter  = regexp.MustCompile(`(?i)^(?:sep|sep.|september|q3|jul)?\s+(\d{4})\s*$`)
	reDecQuarter  = regexp.MustCompile(`(?i)^(?:dec|dec.|december|q4|oct)?\s+(\d{4})\s*$`)
	reQuarterPost = regexp.MustCompile(`(?i)^(\d{4})\s*q([1-4])\s*$`)
)

var dateFormats = []string{"_2 Jan 2006", "_2 January 2006", "_2 Jan, 2006", "_2 January, 2006", "January _2 2006", "Jan _2 2006", "Jan _2, 2006"}

var defaultParser = Parser{}

// Parse uses heuristics to parse s into the highest precision date available using the default parser
// which does not set a calendar and only handles English.
// An Unknown date is returned for any string that does not contain a detectable date.
func Parse(s string) (Date, error) {
	return defaultParser.Parse(s)
}

// A Parser converts strings into dates
type Parser struct {
	// TODO: options such as calendar and language
}

// Parse uses heuristics to parse s into the highest precision date available.
// An Unknown date is returned for any string that does not contain a detectable date.
func (p *Parser) Parse(s string) (Date, error) {
	for _, f := range dateFormats {
		if t, err := time.Parse(f, s); err == nil {
			return &Precise{
				Y: t.Year(),
				M: int(t.Month()),
				D: t.Day(),
			}, nil
		}
	}

	if reYear.MatchString(s) {
		y, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		return &Year{
			Y: y,
		}, nil
	}

	m := reBeforeYear.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &BeforeYear{
			Y: y,
		}, nil

	}

	m = reAfterYear.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &AfterYear{
			Y: y,
		}, nil

	}

	m = reAboutYear.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &AboutYear{
			Y: y,
		}, nil

	}

	m = reMarQuarter.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &YearQuarter{
			Y: y,
			Q: 1,
		}, nil

	}

	m = reJunQuarter.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &YearQuarter{
			Y: y,
			Q: 2,
		}, nil

	}

	m = reSepQuarter.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &YearQuarter{
			Y: y,
			Q: 3,
		}, nil

	}

	m = reDecQuarter.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		return &YearQuarter{
			Y: y,
			Q: 4,
		}, nil

	}

	m = reQuarterPost.FindStringSubmatch(s)
	if len(m) > 1 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		q, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		return &YearQuarter{
			Y: y,
			Q: q,
		}, nil

	}

	return &Unknown{Text: s}, nil
}
