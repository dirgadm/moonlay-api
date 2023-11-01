package ehttp

import "net/http"

type FormatResponse struct {
	Code    int               `json:"code"`
	Status  string            `json:"status,omitempty"`
	Message interface{}       `json:"message,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
	Total   int64             `json:"total,omitempty"`
	Page    int               `json:"page,omitempty"`
	PerPage int               `json:"per_page,omitempty"`
}

func (r *FormatResponse) SetMessage(message string) *FormatResponse {
	r.Status = "success"
	r.Code = http.StatusOK
	r.Message = message

	return r
}

// SetData fill data and total into response formater.
func (r *FormatResponse) SetData(d interface{}) *FormatResponse {
	r.Status = "success"
	r.Code = http.StatusOK
	r.Data = &d

	return r
}

// SetDataList fill data list into response formater.
func (r *FormatResponse) SetDataList(d interface{}, total int64, page int, perPage int) *FormatResponse {
	r.Status = "success"
	r.Code = http.StatusOK
	r.Data = &d

	if total > 0 {
		r.Total = total
		r.PerPage = perPage
		r.Page = page
	}

	return r
}
