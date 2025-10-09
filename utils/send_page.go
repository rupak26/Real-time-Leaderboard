package utils

import (
	"net/http"
)

type PaginatedData struct {
    DataList    any `json:"datalist"`
	Pagination  Pagination
}

type Pagination struct {
	Limit      int64             `json:"limit"`
} 


func SendPage(w http.ResponseWriter , data any  , limit int64)  {
     paginatedData := PaginatedData{
		DataList: data,
		Pagination: Pagination{
			Limit: limit,
		},
	 }
	 WriteResponse(w , http.StatusOK , paginatedData)
}