package model

import "time"

type Article struct {
	ArticleId int64 `db:"id"`
	CategoryId int64 `db:"category_id"`
	Title string `db:"title"`
	ViewCount uint16 `db:"view_count"`
	CommentCount uint16 `db:"comment_count"`
	UserName string `db:"username"`
	Status int `db:"status"`
	Summary string `db:"summary"`
	Create_time time.Time `db:"create_time"`
}

type ArticleDetail struct {
	Article
	Content  string `db:"content"`
	Category
}

type ArticleList struct {
	Article
	Category
}
