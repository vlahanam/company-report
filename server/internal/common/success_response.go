package common

import (
	"fmt"
	"strings"
)

type SuccessResponse struct {
	Status     string      `json:"status"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Limit      int     `json:"limit"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
	PrevCursor *string `json:"prev_cursor,omitempty"`
	TotalItems *int64  `json:"total_items,omitempty"`
}

func NewSuccessResponse(code, message string) *SuccessResponse {
	return &SuccessResponse{
		Status:  "SUCCESS",
		Code:    code,
		Message: message,
	}
}

func NewCreatedResponse(resource, message string) *SuccessResponse {
	return NewSuccessResponse(
		fmt.Sprintf("CREATE_%s_SUCCESSFULLY", strings.ToUpper(resource)),
		message,
	)
}

func NewUpdatedResponse(resource, message string) *SuccessResponse {
	return NewSuccessResponse(
		fmt.Sprintf("UPDATE_%s_SUCCESSFULLY", strings.ToUpper(resource)),
		message,
	)
}

func NewGetResponse(resource, message string) *SuccessResponse {
	return NewSuccessResponse(
		fmt.Sprintf("GET_%s_SUCCESSFULLY", strings.ToUpper(resource)),
		message,
	)
}

func NewDeletedResponse(resource, message string) *SuccessResponse {
	return NewSuccessResponse(
		fmt.Sprintf("DELETE_%s_SUCCESSFULLY", strings.ToUpper(resource)),
		message,
	)
}

func (r *SuccessResponse) WithData(data interface{}) *SuccessResponse {
	r.Data = data
	return r
}

func (r *SuccessResponse) WithPagination(limit int, nextCursor, prevCursor string, hasMore bool) *SuccessResponse {
	p := &Pagination{
		Limit:   limit,
		HasMore: hasMore,
	}

	if nextCursor != "" {
		p.NextCursor = &nextCursor
	}
	if prevCursor != "" {
		p.PrevCursor = &prevCursor
	}

	r.Pagination = p
	return r
}

func (r *SuccessResponse) WithTotalItems(total int64) *SuccessResponse {
	if r.Pagination == nil {
		r.Pagination = &Pagination{}
	}
	r.Pagination.TotalItems = &total
	return r
}
