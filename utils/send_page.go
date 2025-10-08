package utils

import (
	"net/http"
)

type PaginatedData struct {
    DataList    any `json:"datalist"`
	Pagination  Pagination
}

type Pagination struct {
	Page       int64             `json:"page"`
	Limit      int64             `json:"limit"`
} 


func SendPage(w http.ResponseWriter , data any  , page , limit int64)  {
     paginatedData := PaginatedData{
		DataList: data,
		Pagination: Pagination{
			Page: page,
			Limit: limit,
		},
	 }
	 WriteResponse(w , http.StatusOK , paginatedData)
}