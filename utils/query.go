package utils

import (
	"belajariah-main-service/model"
	"fmt"
)

func GetOrderHandler(orders []model.Order) string {

	var orderQuery string
	for _, value := range orders {

		newOrder := fmt.Sprintf(`%s %s`, value.Field, value.Dir)
		if len(orderQuery) > 0 {
			orderQuery += ", " + newOrder
		} else {
			orderQuery = "ORDER BY " + newOrder

		}
	}

	return orderQuery
}

func GetFilterOrderHandler(defaultFilter, defaultOrder string, query model.Query) string {
	finalOrder := defaultOrder
	finalFilter := defaultFilter
	if len(query.Orders) > 0 {
		finalOrder = GetOrderHandler(query.Orders)
	}

	if len(query.Filters) > 0 {
		// if default filter is not empty than join the filter. else create new filter script
		if len(defaultFilter) > 0 {
			finalFilter = fmt.Sprintf(`WHERE %s AND (%s)`, defaultFilter, GetFilterHandler(query.Filters))
		} else {
			finalFilter = fmt.Sprintf(`WHERE %s `, GetFilterHandler(query.Filters))

		}
	} else {
		finalFilter = fmt.Sprintf(`WHERE %s`, finalFilter)
	}
	var finalResult = fmt.Sprintf(`%s %s`, finalFilter, finalOrder)

	return finalResult
}
