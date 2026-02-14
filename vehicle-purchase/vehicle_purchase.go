// Package purchase helps you prepare to buy a vehicle.
package purchase


// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	choice = choice + " is clearly the better choice."
	return choice
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	resellPrice := float64(originalPrice)
	if age < 3 {
		resellPrice = (resellPrice * 80) / 100
	} else if age >= 3 && age < 10 {
		resellPrice = (resellPrice * 70) / 100
	} else if age >= 10 {
		resellPrice = (resellPrice * 50) / 100
	}
	return resellPrice
}
