package gdate

import (
	"regexp"
	"strconv"
	"time"
)

const (
	janAlts = `jan|jan\.|january|janry`
	febAlts = `feb|feb\.|february|febry`
	marAlts = `mar|mar\.|march`
	aprAlts = `apr|apr\.|april`
	mayAlts = `may`
	junAlts = `jun|jun\.|june`
	julAlts = `jul|jul\.|july`
	augAlts = `aug|aug\.|august`
	sepAlts = `sep|sep\.|september`
	octAlts = `oct|oct\.|october`
	novAlts = `nov|nov\.|november`
	decAlts = `dec|dec\.|december`
)

var (
	reYear        = regexp.MustCompile(`^\d{4}$`)
	reBeforeYear  = regexp.MustCompile(`(?i)^bef(?:.|ore)?\s+(\d{4})\s*$`)
	reAfterYear   = regexp.MustCompile(`(?i)^aft(?:.|er)?\s+(\d{4})\s*$`)
	reAboutYear   = regexp.MustCompile(`(?i)^(?:abt|abt.|about)?\s+(\d{4})\s*$`)
	reQuarterPost = regexp.MustCompile(`(?i)^(\d{4})\s*q([1-4])\s*$`)
	reYearRange   = regexp.MustCompile(`^(\d{4})-(\d{4})$`)

	reQuarter = [4]*regexp.Regexp{
		regexp.MustCompile(`(?i)^(?:` + marAlts + `|q1|` + janAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + junAlts + `|q2|` + aprAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + sepAlts + `|q3|` + julAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + decAlts + `|q4|` + octAlts + `)?\s+(\d{4})\s*$`),
	}

	reMonthYearYM = regexp.MustCompile(`^(\d{4})-((?:0[1-9]|1[0-2]|[1-9]))$`)
	reMonthYearMY = regexp.MustCompile(`^((?:0[1-9]|1[0-2]|[1-9]))-(\d{4})$`)

	reMonthYearNamed = [12]*regexp.Regexp{
		regexp.MustCompile(`(?i)^(?:` + janAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + febAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + marAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + aprAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + mayAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + junAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + julAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + augAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + sepAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + octAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + novAlts + `)?\s+(\d{4})\s*$`),
		regexp.MustCompile(`(?i)^(?:` + decAlts + `)?\s+(\d{4})\s*$`),
	}
)

var dateFormats = []string{"_2 Jan 2006", "_2 January 2006", "_2 Jan, 2006", "_2 January, 2006", "January _2 2006", "Jan _2 2006", "Jan _2, 2006", "2006-01-02"}

var defaultParser = Parser{}

// Parse uses heuristics to parse s into the highest precision date available using the default parser
// which does not set a calendar and only handles English.
// An Unknown date is returned for any string that does not contain a detectable date.
func Parse(s string) (Date, error) {
	return defaultParser.Parse(s)
}

// A Parser converts strings into dates
type Parser struct {
	// TODO: options such language
	ReckoningLocation ReckoningLocation

	// AssumeGROQuarter controls whether the parse will assume that ambiguous dates consisting of a month and a year,
	// where the month is the start or end of a quarter, refer to the UK General Register Office quarter
	// containing that month, so July 1850 will be parsed as 3rd Quarter, 1850
	AssumeGROQuarter bool
}

// Parse uses heuristics to parse s into the highest precision date available.
// An Unknown date is returned for any string that does not contain a detectable date.
func (p *Parser) Parse(s string) (Date, error) {
	for _, f := range dateFormats {
		if t, err := time.Parse(f, s); err == nil {
			return &Precise{
				C: p.ReckoningLocation.Calendar(t.Year()),
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
			C: p.ReckoningLocation.Calendar(y),
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
			C: p.ReckoningLocation.Calendar(y - 1),
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
			C: p.ReckoningLocation.Calendar(y + 1),
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
			C: p.ReckoningLocation.Calendar(y),
			Y: y,
		}, nil

	}

	if p.AssumeGROQuarter {
		d, err := p.tryParseQuarter(s)
		if err != nil {
			return nil, err
		}
		if d != nil {
			return d, nil
		}

		d, err = p.tryParseMonthYear(s)
		if err != nil {
			return nil, err
		}
		if d != nil {
			return d, nil
		}
	} else {
		d, err := p.tryParseMonthYear(s)
		if err != nil {
			return nil, err
		}
		if d != nil {
			return d, nil
		}

		d, err = p.tryParseQuarter(s)
		if err != nil {
			return nil, err
		}
		if d != nil {
			return d, nil
		}

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
			C: p.ReckoningLocation.Calendar(y),
			Y: y,
			Q: q,
		}, nil

	}

	m = reYearRange.FindStringSubmatch(s)
	if len(m) > 2 {
		lower, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		upper, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		return &YearRange{
			C:     p.ReckoningLocation.Calendar(lower),
			Lower: lower,
			Upper: upper,
		}, nil

	}

	return &Unknown{Text: s}, nil
}

func (p *Parser) tryParseQuarter(s string) (Date, error) {
	for i, re := range reQuarter {
		m := re.FindStringSubmatch(s)
		if len(m) > 1 {
			y, err := strconv.Atoi(m[1])
			if err != nil {
				return nil, err
			}
			return &YearQuarter{
				C: p.ReckoningLocation.Calendar(y),
				Y: y,
				Q: i + 1,
			}, nil

		}

	}

	return nil, nil
}

func (p *Parser) tryParseMonthYear(s string) (Date, error) {
	m := reMonthYearYM.FindStringSubmatch(s)
	if len(m) > 2 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		mo, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		return &MonthYear{
			C: p.ReckoningLocation.Calendar(y),
			Y: y,
			M: mo,
		}, nil
	}
	m = reMonthYearMY.FindStringSubmatch(s)
	if len(m) > 2 {
		mo, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		return &MonthYear{
			C: p.ReckoningLocation.Calendar(y),
			Y: y,
			M: mo,
		}, nil
	}

	for i, re := range reMonthYearNamed {
		m := re.FindStringSubmatch(s)
		if len(m) > 1 {
			y, err := strconv.Atoi(m[1])
			if err != nil {
				return nil, err
			}
			return &MonthYear{
				C: p.ReckoningLocation.Calendar(y),
				Y: y,
				M: i + 1,
			}, nil

		}

	}

	return nil, nil
}
