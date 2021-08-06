package model

type Spec struct {
	SpecsName string `json:"spec_name" db:"specs_name"`
	SpecsValueName string `json:"specs_value_name" db:"specs_value_name"`
	Id int64 `json:"id" db:"id"`
}
