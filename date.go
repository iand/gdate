package gdate

import (
	"fmt"
	"strconv"
)

type Date interface {
	String() string
	SortsBefore(Date) bool
	Occurrence() string
}

// SortsBefore reports whether a should sort before b chronologically
func SortsBefore(a, b Date) bool {
	if a == nil || b == nil {
		return false
	}
	return a.SortsBefore(b)
}

// AsYear returns the date as a Year and true if possible, false if it is not possible to convert.
func AsYear(d Date) (*Year, bool) {
	if yearer, ok := d.(interface{ Year() int }); ok {
		return &Year{Y: yearer.Year()}, true
	}
	return nil, false
}

// AsPrecise returns the date as a precise date and true if possible, false if it is not possible to convert.
func AsPrecise(d Date) (*Precise, bool) {
	if p, ok := d.(*Precise); ok {
		return p, true
	}
	return nil, false
}

// IsUnknown reports whether d is an Unknown date
func IsUnknown(d Date) bool {
	if d == nil {
		return true
	}
	_, ok := d.(*Unknown)
	return ok
}

// Unknown is an unknown date. It sorts after every other type of date.
type Unknown struct {
	Text string
}

func (u *Unknown) String() string {
	if u.Text == "" {
		return "unknown"
	}
	return u.Text
}

func (u *Unknown) Occurrence() string {
	return "on an unknown date"
}

func (u *Unknown) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Unknown:
		return u.Text < td.Text
	}
	return false
}

// Precise is a date with a known year, month and day. No calendar is assumed for this date.
type Precise struct {
	Y int
	M int // 1-12 like go's time package
	D int
}

func (p *Precise) String() string {
	return fmt.Sprintf("%d %s %04d", p.D, shortMonthNames[p.M], p.Y)
}

func (p *Precise) Occurrence() string {
	return fmt.Sprintf("on %d %s, %04d", p.D, shortMonthNames[p.M], p.Y)
}

func (p *Precise) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Precise:
		if p.Y != td.Y {
			return p.Y < td.Y
		}
		if p.M != td.M {
			return p.M < td.M
		}
		if p.D != td.D {
			return p.D < td.D
		}
		return false
	case *Year:
		return p.Y < td.Y
	case *BeforeYear:
		return p.Y < td.Y
	case *AfterYear:
		return p.Y <= td.Y
	case *AboutYear:
		return p.Y < td.Y
	case *YearQuarter:
		if p.Y != td.Y {
			return p.Y < td.Y
		}
		if p.M != ((td.Q-1)*3 + 1) {
			return p.M < ((td.Q-1)*3 + 1)
		}
		// Sorts after the start of the quarter
		return p.D == 0
	case *MonthYear:
		if p.Y != td.Y {
			return p.Y < td.Y
		}
		if p.M != td.M {
			return p.M < td.M
		}
		return p.D == 0
	case *EstimatedYear:
		return p.Y < td.Y
	case *YearRange:
		return p.Y < td.Lower
	case *Unknown:
		return true
	}
	return false
}

func (p *Precise) Year() int {
	return p.Y
}

func (p *Precise) DateInYear(long bool) string {
	if long {
		return fmt.Sprintf("%d %s", p.D, longMonthNames[p.M])
	}
	return fmt.Sprintf("%d %s", p.D, shortMonthNames[p.M])
}

// Year is a date for which only the year is known or a period of time that may span an entire year.
// It sorts before any date with a higher numeric year.
type Year struct {
	Y int
}

func (y *Year) String() string {
	return fmt.Sprintf("%04d", y.Y)
}

func (y *Year) Occurrence() string {
	return fmt.Sprintf("in %04d", y.Y)
}

func (y *Year) SortsBefore(d Date) bool {
	// sorts like it is just before the first day of the year
	p := Precise{Y: y.Y, M: 1, D: 0}
	return p.SortsBefore(d)
}

func (y *Year) Year() int {
	return y.Y
}

// Year is a date for which only the month and year is known or a period of time that may span an entire month.
// It sorts before any date with a higher numeric year.
type MonthYear struct {
	M int
	Y int
}

func (m *MonthYear) String() string {
	return fmt.Sprintf("%s %04d", shortMonthNames[m.M], m.Y)
}

func (m *MonthYear) Occurrence() string {
	return fmt.Sprintf("in %s %04d", shortMonthNames[m.M], m.Y)
}

func (m *MonthYear) SortsBefore(d Date) bool {
	// sorts like it is just before first of the month
	p := Precise{Y: m.Y, M: m.M, D: 0}
	return p.SortsBefore(d)
}

func (m *MonthYear) Year() int {
	return m.Y
}

// BeforeYear represents a date that is before the start of a specific year.
// It sorts before any date with that year.
type BeforeYear struct {
	Y int
}

func (b *BeforeYear) String() string {
	return "bef. " + strconv.Itoa(b.Y)
}

func (b *BeforeYear) Occurrence() string {
	return fmt.Sprintf("before %04d", b.Y)
}

func (b *BeforeYear) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Precise:
		return b.Y <= td.Y
	case *Year:
		return b.Y <= td.Y
	case *BeforeYear:
		return b.Y < td.Y
	case *AfterYear:
		return b.Y <= td.Y
	case *AboutYear:
		return b.Y <= td.Y
	case *YearQuarter:
		return b.Y <= td.Y
	case *EstimatedYear:
		return b.Y <= td.Y
	case *MonthYear:
		return b.Y <= td.Y
	case *YearRange:
		return b.Y < td.Lower
	case *Unknown:
		return true
	}
	return false
}

// AfterYear represents a date that is after the end of a specific year
type AfterYear struct {
	Y int
}

func (a *AfterYear) String() string {
	return "aft. " + strconv.Itoa(a.Y)
}

func (a *AfterYear) Occurrence() string {
	return fmt.Sprintf("after %04d", a.Y)
}

func (a *AfterYear) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Precise:
		return a.Y < td.Y
	case *Year:
		return a.Y < td.Y
	case *BeforeYear:
		return a.Y < td.Y
	case *AfterYear:
		return a.Y < td.Y
	case *AboutYear:
		return a.Y < td.Y
	case *YearQuarter:
		return a.Y < td.Y
	case *EstimatedYear:
		return a.Y < td.Y
	case *MonthYear:
		return a.Y < td.Y
	case *YearRange:
		return a.Y < td.Lower
	case *Unknown:
		return true
	}
	return false
}

// AboutYear represents a date that is near to a specific year
type AboutYear struct {
	Y int
}

func (a *AboutYear) String() string {
	return "abt. " + strconv.Itoa(a.Y)
}

func (a *AboutYear) Occurrence() string {
	return fmt.Sprintf("about %04d", a.Y)
}

func (a *AboutYear) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Precise:
		return a.Y <= td.Y
	case *Year:
		return a.Y < td.Y
	case *BeforeYear:
		return a.Y < td.Y
	case *AfterYear:
		return a.Y <= td.Y
	case *AboutYear:
		return a.Y < td.Y
	case *YearQuarter:
		return a.Y <= td.Y
	case *EstimatedYear:
		return a.Y < td.Y
	case *MonthYear:
		return a.Y < td.Y
	case *YearRange:
		return a.Y < td.Lower
	case *Unknown:
		return true
	}
	return false
}

func (a *AboutYear) Year() int {
	return a.Y
}

// YearQuarter represents quarter of a specific year, based on GRO quarters
// Values of Q correspond to quarters as follows:
// 1 = Jan-Mar, known as MAR QTR
// 2 = Apr-Jun, known as JUN QTR
// 3 = Jul-Sep, known as SEP QTR
// 4 = Oct-Dec, known as DEC QTR
type YearQuarter struct {
	Y int
	Q int
}

func (y *YearQuarter) MonthRange() string {
	switch y.Q {
	case 1:
		return "Jan-Mar"
	case 2:
		return "Apr-Jun"
	case 3:
		return "Jul-Sep"
	case 4:
		return "Oct-Dec"
	}
	return "Unknown quarter"
}

func (y *YearQuarter) String() string {
	return fmt.Sprintf("%s %04d", y.MonthRange(), y.Y)
}

func (y *YearQuarter) Occurrence() string {
	return fmt.Sprintf("in the %s quarter of %04d", y.MonthRange(), y.Y)
}

func (y *YearQuarter) SortsBefore(d Date) bool {
	// sorts like it is just before the first day of the quarter
	p := Precise{Y: y.Y, M: 1 + (y.Q-1)*3, D: 0}
	return p.SortsBefore(d)
}

func (y *YearQuarter) Year() int {
	return y.Y
}

// EstimatedYear represents a date that is estimated to be a specific year
type EstimatedYear struct {
	Y int
}

func (e *EstimatedYear) String() string {
	return "est. " + strconv.Itoa(e.Y)
}

func (e *EstimatedYear) Occurrence() string {
	return fmt.Sprintf("estimated %04d", e.Y)
}

func (e *EstimatedYear) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *Precise:
		return e.Y <= td.Y
	case *Year:
		return e.Y < td.Y
	case *BeforeYear:
		return e.Y < td.Y
	case *AfterYear:
		return e.Y <= td.Y
	case *AboutYear:
		return e.Y < td.Y
	case *YearQuarter:
		return e.Y <= td.Y
	case *EstimatedYear:
		return e.Y < td.Y
	case *MonthYear:
		return e.Y < td.Y
	case *YearRange:
		return e.Y < td.Lower
	case *Unknown:
		return true
	}
	return false
}

func (e *EstimatedYear) Year() int {
	return e.Y
}

var shortMonthNames = []string{
	1:  "Jan",
	2:  "Feb",
	3:  "Mar",
	4:  "Apr",
	5:  "May",
	6:  "Jun",
	7:  "Jul",
	8:  "Aug",
	9:  "Sep",
	10: "Oct",
	11: "Nov",
	12: "Dec",
}

var longMonthNames = []string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

// YearRange represents a date that is within the range of two years, including the upper and lower year.
type YearRange struct {
	Lower int // first year of the range
	Upper int // last year of the range
}

func (y *YearRange) String() string {
	if y.Lower%10 == 0 && (y.Upper-y.Lower == 9 || y.Upper-y.Lower == 99) {
		return fmt.Sprintf("%ds", y.Lower)
	}
	return strconv.Itoa(y.Lower) + "-" + strconv.Itoa(y.Upper)
}

func (y *YearRange) Occurrence() string {
	if y.Lower%10 == 0 && (y.Upper-y.Lower == 9 || y.Upper-y.Lower == 99) {
		return fmt.Sprintf("in the %ds", y.Lower)
	}
	return fmt.Sprintf("between %d and %d", y.Lower, y.Upper)
}

func (y *YearRange) SortsBefore(d Date) bool {
	switch td := d.(type) {
	case *YearRange:
		if y.Lower != td.Lower {
			return y.Lower <= td.Lower
		}
		return y.Upper <= td.Upper
	case *Precise:
		return y.Lower <= td.Y
	case *Year:
		return y.Lower < td.Y
	case *BeforeYear:
		return y.Lower < td.Y
	case *AfterYear:
		return y.Upper <= td.Y
	case *AboutYear:
		return y.Lower < td.Y
	case *YearQuarter:
		return y.Lower <= td.Y
	case *EstimatedYear:
		return y.Lower < td.Y
	case *MonthYear:
		return y.Lower < td.Y
	case *Unknown:
		return true
	}
	return false
}
