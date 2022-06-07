package repository

import (
	"douyinProject/entity"
	"errors"

	"gorm.io/gorm"
)

func DoFavorite(act int64, video_id int64, user_id int64) error {
	curUser := DbUser{Id: user_id}
	curVideo := dbVideo{Id: video_id}

	//检查视频和用户是否存在
	video := dbVideo{}
	result := db.Where(curVideo).Find(&video)
	if result.Error != nil || video.Id != video_id {
		return errors.New("视频不存在")
	}
	user := DbUser{}
	result = db.Where(curUser).Find(&user)
	if result.Error != nil || user.Id != user_id {
		return errors.New("用户不存在")
	}
	curFavorite := video.Favoritecount // 当前点赞数量
	//准备用户点赞列表
	favoriteList := UserFavorite{}
	result = db.Model(&UserFavorite{}).Where(DbUser{Id: user_id}).Preload("Videos").Find(&favoriteList)
	if result.Error != nil || favoriteList.Id == 0 { //若无则创建
		favoriteList = UserFavorite{
			User: curUser,
		}
		db.Create(&favoriteList)
		db.Model(&UserFavorite{}).Where(DbUser{Id: user_id}).Preload("Videos").Find(&favoriteList)
	}
	if favoriteList.Id == 0 {
		return errors.New("操作失败")
	}
	var err error
	if act == 1 { // 点赞, 赞+1
		for _, v := range favoriteList.Videos {
			if v.Id == video_id {
				return nil
			}
		}
		db.Model(&curVideo).Update("Favoritecount", curFavorite+1)
		err = db.Model(&favoriteList).Association("Videos").Append(&curVideo)
	} else { //	取消, 赞-1
		ok := false
		for _, v := range favoriteList.Videos {
			if v.Id == video_id {
				ok = true
				break
			}
		}
		if !ok {
			return nil
		}
		db.Model(&curVideo).Update("Favoritecount", curFavorite-1)
		err = db.Model(&favoriteList).Association("Videos").Delete(&curVideo)
	}

	if err != nil {
		db.Rollback() // 更新失败则回滚
		return errors.New("操作失败")
	}
	return nil
}
func FavoriteList(user_id int64) *[]entity.Video {
	list := UserFavorite{}
	// 获取用户点赞列表
	result := db.Model(&UserFavorite{}).Where(DbUser{Id: user_id}).Preload("Videos").Find(&list)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || list.Id == 0 {
		db.Create(&UserFavorite{User: DbUser{Id: user_id}})
	}

	res := make([]entity.Video, len(list.Videos))
	for i, v := range list.Videos { // 转entity.Video
		res[i] = copyValue(v)
	}
	return &res
}
