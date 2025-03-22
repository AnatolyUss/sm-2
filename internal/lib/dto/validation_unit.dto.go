package dto

type ValidationUnitDto struct {
	Name  string `form:"name" json:"name" binding:"required,min=1,max=255"`
	Value any    `form:"value" json:"value" binding:"required"`
}
