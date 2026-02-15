package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	towtBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		towtBirds = towtBirds + birdsPerDay[i]
	}
	return towtBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week * 7) - 7 // first day of the week
	towtBirdsThisWeek := 0
	for i := startIndex; i < (startIndex + 7); i++ {
		towtBirdsThisWeek = towtBirdsThisWeek + birdsPerDay[i]
	}
	return towtBirdsThisWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedBirdsPerDay := birdsPerDay
	fixDay := true // toggle for days to fix
	for i := 0; i < len(birdsPerDay); i++ {
		if (fixDay) {
			fixedBirdsPerDay[i] = fixedBirdsPerDay[i] + 1
			fixDay = false
		} else {
			// toggle for next iteration
			fixDay = true
		}
	}
	return fixedBirdsPerDay
}
