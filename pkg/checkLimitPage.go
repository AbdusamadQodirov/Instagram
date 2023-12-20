package pkg

func Limit(limit int) int {
	if limit <= 0 {
		limit = 10
	}
	return limit
}

func Page(page int) int {
	if page <= 0 {
		page = 1
	}
	return page
}