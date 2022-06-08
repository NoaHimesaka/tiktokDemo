package repository

import (
	"douyinProject/entity"
	"errors"

	"gorm.io/gorm"
)

func FavoritePreCheck(vid int64, uid int64) (video dbVideo, err error) {
	video, err = GetVideoById(vid)
	if err != nil {
		return
	}
	_, err = GetUserById(uid)
	return
}
func GetFavoriteListByUser(uid int64) (UserFavorite, error) {
	favoriteList := UserFavorite{}
	result := db.Model(&UserFavorite{}).Where("user_id = ?", uid).Preload("Videos").Find(&favoriteList)
	if result.Error != nil || favoriteList.Id == 0 || favoriteList.UserId != uid { //若无则创建
		favoriteList = UserFavorite{
			User: DbUser{Id: uid},
		}
		db.Create(&favoriteList)
		db.Model(&UserFavorite{}).Where("user_id = ?", uid).Preload("Videos").Find(&favoriteList)
		if favoriteList.Id == 0 {
			return favoriteList, errors.New("操作失败")
		}
	}
	return favoriteList, nil
}
func DoFavorite(curVideo *dbVideo, favoriteList *UserFavorite) error {
	for _, v := range favoriteList.Videos {
		if v.Id == curVideo.Id {
			return nil
		}
	}
	db.Model(curVideo).Update("Favoritecount", curVideo.Favoritecount+1)
	err := db.Model(favoriteList).Association("Videos").Append(&dbVideo{Id: curVideo.Id})
	if err != nil {
		return errors.New("操作失败")
	}
	return nil
}
func UnFavorite(curVideo *dbVideo, favoriteList *UserFavorite) error {
	ok := false
	for _, v := range favoriteList.Videos {
		if v.Id == curVideo.Id {
			ok = true
			break
		}
	}
	if !ok {
		return nil
	}
	db.Model(curVideo).Update("Favoritecount", curVideo.Favoritecount-1)
	err := db.Model(favoriteList).Association("Videos").Delete(&dbVideo{Id: curVideo.Id})
	if err != nil {
		return errors.New("操作失败")
	}
	return nil
}
func FavoriteAct(act int64, video_id int64, user_id int64) (err error) {
	//检查视频和用户是否存在
	video, err := FavoritePreCheck(video_id, user_id)
	if err != nil {
		return err
	}
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	//准备用户点赞列表
	favoriteList, err := GetFavoriteListByUser(user_id)
	if err != nil {
		return err
	}
	switch act {
	case 1:
		err = DoFavorite(&video, &favoriteList)
	case 2:
		err = UnFavorite(&video, &favoriteList)
	}
	return
}

//空video
var nullVideo = []dbVideo{}

func FavoriteList(user_id int64) *[]entity.Video {
	curUser := DbUser{Id: user_id}
	list := UserFavorite{}
	// 获取用户点赞列表
	result := db.Model(&UserFavorite{}).Where("user_id = ?", user_id).Preload("Videos").Find(&list)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || list.Id == 0 || list.UserId != user_id {
		db.Create(&UserFavorite{User: curUser})
		db.Model(&UserFavorite{}).Where("user_id = ?", user_id).Preload("Videos").Find(&list)
	}
	if list.Id == 0 || list.UserId != user_id {
		list.Videos = nullVideo
	}
	res := make([]entity.Video, len(list.Videos))
	for i, v := range list.Videos { // 转entity.Video
		res[i] = copyValue(v)
	}
	return &res
}
