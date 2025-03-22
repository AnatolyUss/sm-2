package dto

type RequestDto struct {
	RequestModelBaseDto
	QueryParams []*ValidationUnitDto `form:"query_params" json:"query_params" binding:"dive"`
	Headers     []*ValidationUnitDto `form:"headers" json:"headers" binding:"dive"`
	Body        []*ValidationUnitDto `form:"body" json:"body" binding:"dive"`
}
