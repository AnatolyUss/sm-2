package dto

import "sm-2/main/internal/lib/types"

type RequestModelBaseDto struct {
	Method types.HttpMethod `form:"method" json:"method" binding:"required,oneof=TRACE PUT POST PATCH OPTIONS HEAD GET DELETE CONNECT"`
	Path   string           `form:"path" json:"path" binding:"required,min=1,max=255"`
}
