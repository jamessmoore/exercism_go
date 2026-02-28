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
	return fmt.Sprintf("This is the box containing the number %s", prepared)
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
	var isSupplied int
	return isSupplied
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	var result string
	return result
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	var result string
	return result
}
 