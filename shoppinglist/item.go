package shoppinglist

import "time"

// CartItem represent the structure of each item stored in shopping list
type CartItem struct {
	ItemName    string
	Category    string
	DateEntered time.Time
	User        string
}

func getItemsByCategory(items []CartItem, category string) (ret []CartItem) {
	for _, element := range items {
		if element.Category == category {
			ret = append(ret, element)
		}
	}
	return
}

func getItemsByDateRange(items []CartItem, fromDate time.Time, toDate time.Time) (ret []CartItem) {
	for _, element := range items {
		if element.DateEntered.After(fromDate) && element.DateEntered.Before(toDate) {
			ret = append(ret, element)
		}
	}
	return
}

func getItemsByDate(items []CartItem, date time.Time) (ret []CartItem) {
	return getItemsByDateRange(items, date, date)
}

func getItemsByDateAndCategory(items []CartItem, date time.Time, category string) (ret []CartItem) {
	for _, element := range items {
		if element.DateEntered.Equal(date) && element.Category == category {
			ret = append(ret, element)
		}
	}
	return
}

func getAllCategories(items []CartItem) []string {
	categories := make([]string, 0)
	for _, element := range items {
		categories = append(categories, element.Category)
	}

	return categories
}
