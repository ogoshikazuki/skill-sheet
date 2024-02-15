package repository

import (
	"fmt"
	"strconv"
)

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

func setUpdate(count int, column string) string {
	var query string
	if count > 1 {
		query += ","
	}
	query += fmt.Sprintf(` "%s" = $%d`, column, count)

	return query
}
