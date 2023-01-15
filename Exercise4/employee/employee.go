package employee

import (
	"strconv"
	"time"
)

type Employee struct {
	FirstName   string
	LastName    string
	DateOfBirth string
}

// GreetEmployee returns the employee name with greeting
func GreetEmployee(employee Employee) string {
	if len(employee.FirstName) == 0 && len(employee.LastName) == 0 {
		return "Hello"
	} else if len(employee.FirstName) == 0 {
		return "Hello " + employee.LastName
	} else if len(employee.LastName) == 0 {
		return "Hello " + employee.FirstName
	}
	return "Hello " + employee.FirstName + " " + employee.LastName
}

// FindAge returns the age based on Date of Birth
func FindAge(employee Employee) int {
	dob := employee.DateOfBirth
	m := map[time.Month]int{
		time.January:   1,
		time.February:  2,
		time.March:     3,
		time.April:     4,
		time.May:       5,
		time.June:      6,
		time.July:      7,
		time.August:    8,
		time.September: 9,
		time.October:   10,
		time.November:  11,
		time.December:  12,
	}
	date, err1 := strconv.Atoi(dob[0:2])
	month, err2 := strconv.Atoi(dob[3:5])
	year, err3 := strconv.Atoi(dob[6:10])
	if err1 != nil || err2 != nil || err3 != nil {
		return -1
	}
	age := 0
	currentDate := time.Now().Day()
	currentMonth := m[time.Now().Month()]
	currentYear := time.Now().Year()
	if year > currentYear {
		return -1
	} else if year == currentYear {
		if month > currentMonth {
			return -1
		} else if month == currentMonth {
			if date > currentDate {
				return -1
			}
		}
	}
	age = currentYear - year
	if month > currentMonth {
		age--
		return age
	} else if month == currentMonth {
		if date > currentDate {
			age--
			return age
		}
	}
	return age
}
