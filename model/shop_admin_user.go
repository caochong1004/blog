package model

type ShopAdminUsers struct {
	Id int64 `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Status int `db:"status" json:"status"`
}

type LoginResult struct {
	Token string `json:"token"`
	Username string `json:"username"`
}
