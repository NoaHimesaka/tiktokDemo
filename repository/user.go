package repository

import (
	"douyinProject/entity"
	"douyinProject/utils"
	"errors"

	"gorm.io/gorm"
)

func CheckUserName(username string) error {
	//if(查询结果不存在) 返回error
	result := db.Where("username = ?", username).First(&DbUser{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	//else 返回nil
	return errors.New("用户名已经存在")
}
func GetUserById(uid int64) (DbUser, error) {
	user := DbUser{}
	result := db.Where(&DbUser{Id: uid}).Find(&user)
	if result.Error != nil || user.Id != uid {
		return user, errors.New("用户不存在")
	}
	return user, nil
}
func Register(username string, password string) (int64, string, error) {
	user := DbUser{
		Username:      username,
		Password:      password,
		Followcount:   0,
		Followercount: 0,
	}
	//返回插入后的id,nil
	result := db.Create(&user) // 通过数据的指针来创建
	token, _ := utils.GenerateToken(user.Id)
	return user.Id, token, result.Error
}
func Login(username string, password string) (int64, string, error) {

	user := DbUser{}
	//返回插入后的id,token,nil
	result := db.Where("username = ? And password = ?", username, password).
		First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return -1, "", errors.New("用户名或者密码错误")
	}
	token, _ := utils.GenerateToken(user.Id)
	return user.Id, token, result.Error
}
func GetUserInfo(userId int64) (entity.User, error) {
	curUser := DbUser{}
	result := db.Where("id = ?", userId).First(&curUser)
	return copyUser(curUser), result.Error
}
func GetUserByToken(token string) int64 {
	//由于token用了jwt，直接在jwt payload里放userid
	return utils.GetIdInToken(token)
}
func copyUser(User2 DbUser) entity.User {
	curUser := entity.User{}
	curUser.Name = User2.Username
	curUser.Id = User2.Id
	curUser.FollowCount = User2.Followcount
	curUser.FollowerCount = User2.Followercount
	curUser.IsFollow = false
	return curUser
}
