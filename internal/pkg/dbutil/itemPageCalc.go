package dbutil

import "math"

// ItemPageCalc calculates item numbers needed for page navigation
func ItemPageCalc(perPage, page, itemCount int64) (startNum, totalPages int64) {
	//calculate total page numbers
	totalPages = int64(math.Ceil(float64(itemCount) / float64(perPage)))
	if page < 1 {
		page = 1
	}
	//calculate number to start with
	startNum = (page - 1) * perPage
	return
}
