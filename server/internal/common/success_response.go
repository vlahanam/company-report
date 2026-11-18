package common

type successResponse struct {
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

func NewSuccessResponse(code, message string) *successResponse {
	return &successResponse{
		Status:  "SUCCESS",
		Code:    code,
		Message: message,
	}
}

func (r *successResponse) WithData(data interface{}) *successResponse {
	r.Data = data
	return r
}

func (r *successResponse) WithPagination(limit int, nextCursor, prevCursor string, hasMore bool) *successResponse {
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

func (r *successResponse) WithTotalItems(total int64) *successResponse {
	if r.Pagination == nil {
		r.Pagination = &Pagination{}
	}
	r.Pagination.TotalItems = &total
	return r
}
