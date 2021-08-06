package service

import (
	"bytes"
	"errors"
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Login(username, password string) (r *model.LoginResult, err error)  {
	user, err  := db.GetUserInfo(username)
	if err != nil{
		return
	}
	if user.Username == "" {
		err = errors.New("用户不存在")
		return
	}

	if user.Status != 1 {
		err = errors.New("用户已被禁用，请联系管理员")
		return
	}

	if user.Password != password {
		err = errors.New("密码错误")
	}
	token := RandString(32)
	var bt bytes.Buffer
	bt.WriteString("token:")
	bt.WriteString(token)
	key := bt.String()
	db.RedisConn.Do("SET",key, user.Id,  "EX", "7200")
	r = &model.LoginResult{
		Token: token,
		Username: user.Username,
	}
	return

}


func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}