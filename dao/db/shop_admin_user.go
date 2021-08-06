package db

import "github.com/longjoy/blog/model"

func GetUserInfo(username string) (usi *model.ShopAdminUsers, err error)  {
	usi = &model.ShopAdminUsers{}
	sql := "select id, username, password,status from shop_admin_users where username = ?"
	err = DBWeb.Get(usi, sql, username)
	if err != nil {
		return
	}
	return
}
