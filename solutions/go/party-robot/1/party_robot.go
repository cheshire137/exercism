package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    units := "years"
    if age == 1 {
        units = "year"
    }
	return fmt.Sprintf("Happy birthday %s! You are now %d %s old!", name, age, units)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    directionPrefix := ""
    if direction == "left" || direction == "right" {
        directionPrefix = "on the "
    }
	return fmt.Sprintf("%s\nYou have been assigned to table %0*d. Your table is %s%s, exactly %0.1f meters from here.\nYou will be sitting next to %s.", Welcome(name), 3, table, directionPrefix, direction, distance, neighbor)
}
