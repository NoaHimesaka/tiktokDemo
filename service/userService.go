package service

import (
	"crypto/md5"
	"douyinProject/entity"
	"douyinProject/repository"
	"douyinProject/utils"
	"errors"
	"fmt"
)

//这个方法返回注册失败的异常情况
func ContainsUsername(username string) error {
	err := repository.CheckUserName(username)
	if err != nil {
		return err
	}
	return nil
}

func NotContainsUsername(username string) error {
	err := repository.CheckUserName(username)
	if err == nil {
		return errors.New("不存在该用户名")
	}
	return nil
}

//对用户进行注册
func UserRegister(username string, password string) (int64, string, error) {
	if username == "" {
		return -1, "", errors.New("用户名不能为空")
	}
	if len(password) < 5 {
		return -1, "", errors.New("密码不能少于5个字符")
	}
	passwordByteFlow := md5.Sum(utils.String2Bytes(password))
	md5Password := fmt.Sprintf("%x", passwordByteFlow)
	id, token, err := repository.Register(username, md5Password)
	return id, token, err
}
func UserLogin(username string, password string) (int64, string, error) {
	passwordByteFlow := md5.Sum(utils.String2Bytes(password))
	md5Password := fmt.Sprintf("%x", passwordByteFlow)
	id, token, err := repository.Login(username, md5Password)
	return id, token, err
}

// todo GetuserInfo函数还没写
func GetUserInfo(user_id int64, token string) (entity.User, error) {
	//判断token是否合法
	if !utils.TokenCheck(token) {
		return entity.User{}, errors.New("Invalid Token")
	}
	//
	User, err := repository.GetUserInfo(user_id)
	return User, err
}
