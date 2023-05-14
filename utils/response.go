package utils

type PaginatedResult struct {
    Page         int         `json:"page"`
    PerPage      int         `json:"perPage"`
    TotalRecords int         `json:"totalRecords"`
    TotalPages   int         `json:"totalPages"`
    Items        interface{} `json:"items"`
    Remaining    int         `json:"remaining"`
}