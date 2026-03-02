package sorting

import (
	"strconv"
	"fmt"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	prepared := strconv.FormatFloat(f, 'f', 1, 64)
	return fmt.Sprintf("This is the number %s", prepared)
}
 
type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	number := nb.Number()
	prepared := strconv.FormatFloat(float64(number), 'f', 1, 64)
	return fmt.Sprintf("This is a box containing the number %s", prepared)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
    if ok {
		if sc, err := strconv.Atoi(fnb.Value()); err == nil {
		// fmt.Printf("%T, %v", sc, sc)
		return sc
		}
	}
	// println(fnb.Value())
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := 0.0
	_, ok := fnb.(FancyNumber)
	// fmt.Printf("%T, %v", sc, sc)
    if ok {
		number = float64(ExtractFancyNumber(fnb))
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", number)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	var result string
	switch val := i.(type) {
	case int:
		result = DescribeNumber(float64(val))
	case float64:
		result = DescribeNumber(val)
	default:
		result = "Return to sender"
	}
	return result
}
 