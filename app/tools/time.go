package tools

import (
	"strconv"
	"time"
)

func GetToday() string {

	intyear, tmonth, intday := time.Now().Date()
	intmonth := int(tmonth)

	year := strconv.Itoa(intyear)
	month := strconv.Itoa(intmonth)
	day := strconv.Itoa(intday)

	return year + "-" + month + "-" + day
}

func GetMonth() int {
	_, month, _ := time.Now().Date()

	return int(month)
}

func GetYear() int {
	year, _, _ := time.Now().Date()

	return year
}
