package service

import (
	"douyinProject/entity"
	"douyinProject/repository"
	"douyinProject/utils"
	"errors"
)

func FavoriteAction(user_id string, token string, video_id string, action_type string) (int64, error) {
	// 转换并校验参数
	uid := utils.Int64(user_id, func() int64 {
		return repository.GetUserByToken(token)
	})
	vid := utils.Int64(video_id, func() int64 {
		return -1
	})
	if vid == -1 {
		return 0, errors.New("视频不存在")
	}
	action := utils.Int64(action_type, func() int64 {
		return 0
	})
	if action == 0 {
		return 0, errors.New("未知行为")
	}
	return action, repository.FavoriteAct(action, vid, uid)
}
func GetFavorite(user_id string, token string) (*[]entity.Video, error) {
	uid := utils.Int64(user_id, func() int64 {
		return repository.GetUserByToken(token)
	})
	videos := repository.FavoriteList(uid)
	return videos, nil
}
