// Package cards 
package cards

var deck = []int{1,2,3,4,5,6,7,8,9,10}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	herFavCards := []int{deck[1],deck[5],deck[8]}
	return herFavCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < len(slice) && index > -1 {
		return slice[index]
	} else {
		return -1
	}
}
// card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
// card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index > -1 {
		slice[index] = value
		return slice
	} else {
		slice = append(slice, value)
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var new_slice []int
	new_slice = append(new_slice, values...)
	new_slice = append(new_slice, slice...)
	return new_slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index > -1 {
		newSlice := append(slice[:index], slice[index+1:]...)
		return newSlice
	} else {
		return slice
	}
}