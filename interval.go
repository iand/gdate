package gdate

import (
	"fmt"
	"time"
)

type Interval interface {
	Precise() string
	Rough() string
}

func IntervalBetween(a, b Date) Interval {
	if IsUnknown(a) || IsUnknown(b) {
		return &UnknownInterval{}
	}
	ap, aok := AsPrecise(a)
	bp, bok := AsPrecise(b)
	if aok && bok {
		in := &PreciseInterval{
			Y: bp.Y - ap.Y,
			M: bp.M - ap.M,
			D: bp.D - ap.D,
		}

		if in.D < 0 {
			t := time.Date(ap.Y, time.Month(ap.M), 32, 0, 0, 0, 0, time.UTC)
			in.D += 32 - t.Day()
			in.M--
		}
		if in.M < 0 {
			in.M += 12
			in.Y--
		}

		return in
	}

	ay, aok := AsYear(a)
	by, bok := AsYear(b)
	if aok && bok {
		diff := by.Y - ay.Y
		if diff < 0 {
			diff = -diff
		}
		return &YearsInterval{Y: diff}
	}

	return &UnknownInterval{}
}

// IsUnknownInterval reports whether in is an unknown interval
func IsUnknownInterval(in Interval) bool {
	if in == nil {
		return true
	}
	_, ok := in.(*UnknownInterval)
	return ok
}

// AsYearsInterval returns the interval as a YearsInterval and true if possible, false if it is not possible to convert.
func AsYearsInterval(in Interval) (*YearsInterval, bool) {
	if yi, ok := in.(*YearsInterval); ok {
		return yi, true
	}
	if yearer, ok := in.(interface{ Years() int }); ok {
		return &YearsInterval{Y: yearer.Years()}, true
	}
	return nil, false
}

// AsPreciseInterval returns the interval as a PreciseInterval and true if possible, false if it is not possible to convert.
func AsPreciseInterval(in Interval) (*PreciseInterval, bool) {
	if pi, ok := in.(*PreciseInterval); ok {
		return pi, true
	}
	if ymder, ok := in.(interface{ YMD() (int, int, int) }); ok {
		y, m, d := ymder.YMD()
		return &PreciseInterval{
			Y: y,
			M: m,
			D: d,
		}, true
	}
	return nil, false
}

type PreciseInterval struct {
	Y, M, D int
}

var _ Interval = (*PreciseInterval)(nil)

func (p *PreciseInterval) Precise() string {
	var str string
	if p.Y > 0 {
		str += pluralise(p.Y, "year")
	}
	if p.M > 0 {
		if p.Y > 0 {
			if p.D == 0 {
				str += " and "
			} else {
				str += ", "
			}
		}
		str += pluralise(p.M, "month")
	}
	if p.D > 0 {
		if p.Y > 0 || p.M > 0 {
			str += " and "
		}
		str += pluralise(p.D, "day")
	}

	return str
}

func (p *PreciseInterval) Years() int {
	return p.Y
}

func (p *PreciseInterval) Months() int {
	return p.Y*12 + p.M
}

func (p *PreciseInterval) ApproxDays() int {
	return p.Y*365 + p.M*30 + p.D
}

func (p *PreciseInterval) YMD() (int, int, int) {
	return p.Y, p.M, p.D
}

func (p *PreciseInterval) Rough() string {
	if p.Y > 0 {
		if p.M > 10 {
			return "nearly " + pluralise(p.Y+1, "year")
		}
		return pluralise(p.Y, "year")
	}
	if p.M > 0 {
		if p.D > 27 {
			return "nearly " + pluralise(p.M+1, "month")
		}
		return pluralise(p.M+1, "month")
	}
	return pluralise(p.D, "day")
}

func pluralise(n int, stem string) string {
	var suffix string
	if n != 1 {
		suffix = "s"
	}
	return fmt.Sprintf("%d %s%s", n, stem, suffix)
}

type UnknownInterval struct{}

var _ Interval = (*UnknownInterval)(nil)

func (i *UnknownInterval) Precise() string {
	return "unknown"
}

func (i *UnknownInterval) Rough() string {
	return "unknown"
}

// YearsInterval represents an interval of time measured in whole years
type YearsInterval struct {
	Y int
}

var _ Interval = (*YearsInterval)(nil)

func (i *YearsInterval) Precise() string {
	return pluralise(i.Y, "year")
}

func (i *YearsInterval) Rough() string {
	return pluralise(i.Y, "year")
}

func (p *YearsInterval) Years() int {
	return p.Y
}

// AboutYearsInterval represents an interval of time measured in an uncertain number of whole years
type AboutYearsInterval struct {
	Y int
}

var _ Interval = (*AboutYearsInterval)(nil)

func (i *AboutYearsInterval) Precise() string {
	return "about " + pluralise(i.Y, "year")
}

func (i *AboutYearsInterval) Rough() string {
	return "about " + pluralise(i.Y, "year")
}

func (p *AboutYearsInterval) Years() int {
	return p.Y
}
