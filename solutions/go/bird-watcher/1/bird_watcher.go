package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for i := 0; i < len(birdsPerDay); i++ {
        total += birdsPerDay[i]
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    if week < 1 {
        return 0
    }
    startIdx := (week - 1) * 7
    if startIdx >= len(birdsPerDay) {
        return 0
    }
    endIdx := startIdx + 7
    return TotalBirdCount(birdsPerDay[startIdx:endIdx])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	corrected := birdsPerDay[:]
    if len(corrected) >= 1 {
    	corrected[0]++
    }
    for i := 2; i < len(corrected); i += 2 {
        corrected[i]++
    }
    return corrected
}
