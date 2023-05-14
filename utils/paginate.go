package utils

import (
	"math"
	"net/http"

	"gorm.io/gorm"
)

// Paginate performs pagination on the given database query, based on the page and per_page parameters in the provided HTTP request.
// It returns a PaginatedResult struct containing the paginated items, as well as metadata such as the total number of records, total number of pages, and the number of remaining items.
// The out parameter should be a pointer to a slice that will hold the paginated items.
// If any errors occur during pagination, an error is returned.
func Paginate(r *http.Request, db *gorm.DB, out interface{}) (PaginatedResult, error) {

	page, perPage := GetPaginationParams(r, 1, 10)

    var result PaginatedResult

    // calculate the offset based on the current page and number of items per page
    offset := (page - 1) * perPage

    // get the total number of records
    var count int64
    if err := db.Model(out).Count(&count).Error; err != nil {
        return result, err
    }

    // get the paginated items
    if err := db.Offset(offset).Limit(perPage).Find(out).Error; err != nil {
        return result, err
    }

    // calculate the total number of pages based on the total number of records and items per page
    totalPages := int(math.Ceil(float64(count) / float64(perPage)))

    // calculate the number of remaining items
    remaining := int(count) - (page * perPage)

    result = PaginatedResult{
        Page:         page,
        PerPage:      perPage,
        TotalRecords: int(count),
        TotalPages:   totalPages,
        Items:        out,
        Remaining:    remaining,
    }

    return result, nil
}