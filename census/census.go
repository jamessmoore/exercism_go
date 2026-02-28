// Package census simulates a system used to collect census data.
package census

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident {
		Name: name,  
		Age: age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// check for nil Resident
	if r == nil {
		return false
	}
	// check for name
	if r.Name == "" {
		return false
	}
	
	if r.Address == nil || len(r.Address) == 0 {
		return false
	}

	// check for street in address
	for k, v := range r.Address {
		if k != "street" {
			return false
		}
		if v == "" {
			return false
		}

	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	fmt.Println(r)
	r.Name = ""
	r.Age = 0
	r.Address = nil
	//fmt.Println(r)
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count ++
		}
	}
	return count
}
