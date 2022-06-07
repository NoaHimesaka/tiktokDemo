package repository

import (
	"douyinProject/entity"
	"fmt"
	"time"
)

type Com struct {
	Id         int64
	Userid     int64
	Video_id   int
	Message    string
	CreateData string
}
type Result struct {
	Id          int64       `json:"id"`
	User        entity.User `json:"user"`
	Content     string      `json:"content"`
	Create_data string      `json:"create_data"`
}

func (Com) TableName() string {
	return "dyuploadcomment"
}
func Uploadcomment(a entity.Comment, videoId int) Com {
	comment := Com{
		Userid:     a.Id,
		Video_id:   videoId,
		Message:    a.Content,
		CreateData: time.Now().String(),
	}
	err := db.AutoMigrate(&Com{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Create(&comment)
	return comment
}
func DeleteComment(videoId int) {
	db.Where("Id = ?", videoId).Delete(&Com{})
}
func Get(videoId int) []Result {
	var comm []Com
	db.Where("video_id = ?", videoId).Find(&comm)
	var res []Result
	for _, j := range comm {
		user1, _ := GetUserInfo(j.Userid)
		res = append(res, Result{
			Id:          j.Id,
			User:        user1,
			Content:     j.Message,
			Create_data: j.CreateData,
		})
	}
	return res
}
