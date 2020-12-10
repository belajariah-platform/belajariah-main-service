package utils

import (
	"belajariah-main-service/model"
	"strings"
)

func GetFilterHandler(filters []model.Filter) string {
	var filterQuery, and, or string
	for _, value := range filters {
		splitter := "|"
		if strings.Contains(value.Type, splitter) {
			var fields []string
			var types []string
			var values []string
			fields = strings.Split(value.Field, splitter)
			values = strings.Split(value.Value, splitter)
			types = strings.Split(value.Type, splitter)
			for index, val := range values {
				if types[index] == "text" {
					filterQuery = filterQuery + or + " lower(" + fields[index] + `) like lower('%` + val + "%') "
					or = " OR "
				} else if types[index] == "date" {
					filterQuery = filterQuery + or + fields[index] + `::DATE='` + val + "'"
					or = " OR "
				} else {
					filterQuery = filterQuery + or + fields[index] + `=` + val
					or = " OR "
				}
			}

		} else if value.Type == "text" {
			filterQuery = filterQuery + and + " lower(" + value.Field + `) like lower('%` + value.Value + "%') "
			and = " AND "
		} else if value.Type == "date" {
			filterQuery = filterQuery + and + value.Field + `::DATE='` + value.Value + "'"
			and = " AND "
		} else {
			filterQuery = filterQuery + and + value.Field + `=` + value.Value
			and = " AND "
		}
	}

	if len(filterQuery) > 0 {
		filterQuery = "AND " + filterQuery
	}
	return filterQuery
}
