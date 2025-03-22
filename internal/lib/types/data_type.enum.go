package types

type DataType string

type DataTypeEnum struct {
	Int       DataType
	String    DataType
	Boolean   DataType
	List      DataType
	Date      DataType
	Email     DataType
	UUID      DataType
	AuthToken DataType
}

func InitializeDataTypeEnum() DataTypeEnum {
	return DataTypeEnum{
		Int:       "Int",
		String:    "String",
		Boolean:   "Boolean",
		List:      "List",
		Date:      "Date",
		Email:     "Email",
		UUID:      "UUID",
		AuthToken: "Auth-Token",
	}
}
