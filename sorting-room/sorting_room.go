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
	fn := 0
	_, ok := fnb.(FancyNumber)
    if ok {
		if sc, err := strconv.Atoi(fnb.Value()); err == nil {
		// fmt.Printf("%T, %v", sc, sc)
		fn = sc
		}
	}
	// println(fnb.Value())
	return fn
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	// println(fnb.Value())
	number := 0.0
	sc, ok := fnb.(FancyNumber)
    if ok {
		// if _, err := strconv.Atoi(fnb.Value()); err == nil {
		fmt.Printf("%T, %v", sc, sc)
		println(fnb.Value())
		// number = ExtractFancyNumber(sc)
		// }
	}
	prepared := strconv.FormatFloat(float64(number), 'f', 1, 64)
	return fmt.Sprintf("This is a fancy box containing the number %s", prepared)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	var result string
	return result
}
 