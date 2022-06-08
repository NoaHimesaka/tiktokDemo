package repository

import (
	"douyinProject/entity"
	"errors"
)

//普通视频流推送
func DefalutVideo(latest_time int64) (*[]entity.Video, int64, error) {
	res := []dbVideo{} //需要返给前端的video list
	//fmt.Println(latest_time)
	result := db.Where("createtime < ?", latest_time).Order("createtime desc").Limit(10).Find(&res)
	//dao怎么和vo对应？ copy函数

	curResult := make([]entity.Video, len(res), 10)
	//curResult := []entity.Video{}
	for i := 0; i < len(res); i++ {
		curResult[i] = copyValue(res[i])
	}

	if len(curResult) == 0 {
		return nil, -1, errors.New("no more video")
	}
	sv := dbVideo{}
	result.Last(&sv)
	return &curResult, sv.Created, result.Error
}
func copyValue(video2 dbVideo) entity.Video {
	videoVo := entity.Video{}
	//需要根据authorid查author信息
	curUser := DbUser{Id: video2.Authorid}
	db.First(&curUser)

	videoVo.Id = video2.Id
	videoVo.Created = video2.Created
	videoVo.Author = copyUser(curUser)
	videoVo.CommentCount = video2.Commentcount
	videoVo.CoverUrl = video2.Coverurl
	videoVo.FavoriteCount = video2.Favoritecount
	videoVo.IsFavorite = false
	videoVo.PlayUrl = video2.Playurl
	videoVo.Title = video2.Title

	return videoVo
}
func PublishVideo(authorid int64, playurl string, coverurl string, title string, timeStamp int64) error {
	curvideo := dbVideo{
		Authorid:      authorid,
		Playurl:       playurl,
		Coverurl:      coverurl,
		Favoritecount: 0,
		Commentcount:  0,
		Title:         title,
		Created:       timeStamp,
	}
	result := db.Create(&curvideo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func PublishList(authorid int64) (*[]entity.Video, error) {
	res := []dbVideo{}
	result := db.Where("authorid = ?", authorid).Order("createtime desc").Find(&res)
	curResult := make([]entity.Video, len(res))
	for i := 0; i < len(res); i++ {
		curResult[i] = copyValue(res[i])
	}
	if len(curResult) == 0 {
		return nil, errors.New("no more video")
	}
	return &curResult, result.Error
}

func GetVideoById(vid int64) (dbVideo, error) {
	video := dbVideo{}
	result := db.Where(&dbVideo{Id: vid}).Find(&video)
	if result.Error != nil || video.Id != vid {
		return video, errors.New("视频不存在")
	}
	return video, nil
}
