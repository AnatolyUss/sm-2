package dto

import "sm-2/main/internal/lib/types"

type ValidationUnitTemplateDto struct {
	Name     string           `form:"name" json:"name" binding:"required,min=1,max=255"`
	Required *bool            `form:"required" json:"required" binding:"required"`
	Types    []types.DataType `form:"types" json:"types" binding:"required,dive,oneof=Auth-Token UUID Email Date List Boolean String Int"`
}
