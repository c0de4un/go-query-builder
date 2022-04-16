package crm

type where struct {
	column   string
	operator string
	value    string
}

func newWhere(column string, operator string, value string) *where {
	var result where

	result.column = column
	result.operator = operator
	result.value = value

	return &result
}

func (predicate *where) compileWhere() string {
	return predicate.column + " " + predicate.operator + " " + predicate.value
}
