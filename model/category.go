package model

type Category struct {
	CategoryId int64 `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo int64 `db:"category_no"`
}
