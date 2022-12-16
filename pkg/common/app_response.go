package common

type appResponse struct {
	Data   any `json:"data"`
	Paging any `json:"paging,omitempty"`
}

func NewSuccessResponse(Data, paging any) *appResponse {
	return &appResponse{Data: Data, Paging: paging}
}

func Response(data any) *appResponse {
	return NewSuccessResponse(data, nil)
}

func ResponseWithPaging(data any, paging any) *appResponse {
	return NewSuccessResponse(data, paging)
}
