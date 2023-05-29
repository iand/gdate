package gdate

import "strconv"

type Calendar int

// The JulianDay is the count of days elapsed since the beginning of the Julian period
func (c Calendar) JulianDay(y, m, d int) int {
	switch c {
	case Gregorian:
		if m == 1 || m == 2 {
			y--
			m += 12
		}
		a := y / 100
		b := a / 4
		c := 2 - a + b
		e := int(365.25 * float64(y+4716))
		f := int(30.6001 * float64(m+1))
		return int(float64(c+d+e+f) - 1524.5)
	case Julian:
		return 367*y - (7*(y+5001+(m-9)/7))/4 + (275*m)/9 + d + 1729777
	case Julian25Mar:
		if m == 1 || m == 2 || (m == 3 && d < 25) {
			y--
			m += 12
		}

		return 367*y - (7*(y+5001+(m-9)/7))/4 + (275*m)/9 + d + 1729777

	default:
		panic("unsupported calendar: " + strconv.Itoa(int(c)))
	}
}

func (c Calendar) String() string {
	switch c {
	case Gregorian:
		return "Gregorian"
	case Julian:
		return "Julian"
	case Julian25Mar:
		return "Julian, year starts on 25 Mar"
	default:
		return "unknown calendar (" + strconv.Itoa(int(c)) + ")"

	}
}

const (
	Gregorian   Calendar = 0
	Julian      Calendar = 1 // Julian calendar with the first day of the year being 1 Jan
	Julian25Mar Calendar = 2 // Julian calendar with the first day of the year being 25 Mar
)

// The ReckoningLocation is the location used to determine the reckoning of the calendar which
// determines the date on which the first day of the year changed to 1 Jan and the date on
// which the calendar changed from Julian to Gregorian
type ReckoningLocation int

const (
	// TODO: non-English reckonings
	EnglandAndWales ReckoningLocation = 0
	Scotland        ReckoningLocation = 1
	Ireland         ReckoningLocation = 2
)

// StartOfYear returns calendar in use for the year specified.
func (r ReckoningLocation) Calendar(y int) Calendar {
	switch r {
	case EnglandAndWales, Ireland:
		if y < 1752 {
			// note that 2 Sep 1752 is immediately followed by 14 Sep 1752
			return Julian25Mar
		}
		return Gregorian
	case Scotland:
		if y < 1600 {
			return Julian25Mar
		} else if y < 1752 {
			return Julian
		}
		return Gregorian
	default:
		panic("unsupported reckoning location: " + strconv.Itoa(int(r)))
	}
}
