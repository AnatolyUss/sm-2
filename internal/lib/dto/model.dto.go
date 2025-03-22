package dto

type ValidationUnitTemplateName string

type RequiredFieldChecked bool

type ModelDto struct {
	RequestModelBaseDto
	QueryParams               []*ValidationUnitTemplateDto `form:"query_params" json:"query_params" binding:"dive"`
	Headers                   []*ValidationUnitTemplateDto `form:"headers" json:"headers" binding:"dive"`
	Body                      []*ValidationUnitTemplateDto `form:"body" json:"body" binding:"dive"`
	GroupsToNamesUnitsMap     map[string]map[ValidationUnitTemplateName]*ValidationUnitTemplateDto
	GroupsToRequiredFieldsMap map[string]map[ValidationUnitTemplateName]RequiredFieldChecked
}
