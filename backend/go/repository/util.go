package repository

import "strconv"

func makePlaceholders(count int) string {
	var placeholders string
	for i := 1; i <= count; i++ {
		placeholders += "$" + strconv.Itoa(i)
		if i < count {
			placeholders += ","
		}
	}

	return placeholders
}
