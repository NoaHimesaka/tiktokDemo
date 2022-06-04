package service

import (
	"douyinProject/repository"
	"errors"
	"fmt"
	"time"
)
import "douyinProject/utils"
import "douyinProject/entity"

//service层成功下返回的信息
type feedVideo struct {
	Latest_time int64
	Token       string
	VideoList   *[]entity.Video
	NextTime    int64
}
type PublishList struct {
	Token     string
	User_id   int64
	VideoList *[]entity.Video
}

//根据输入的token和last-time生成一个实例
func NewVideoInstance(lastestTime string, token string) *feedVideo {
	return &feedVideo{
		Latest_time: utils.Int64(lastestTime, func() int64 {
			return time.Now().Unix()
		}),
		Token: token,
	}
}
func (v *feedVideo) Do() error {
	//根据lasttime得到videolist
	videoList, nextTime, err := repository.DefalutVideo(v.Latest_time)
	if err != nil {
		return err
	}
	v.NextTime = nextTime
	v.VideoList = videoList
	return nil
}
func VideoPublish(token string, title string, data []byte, filename string) error {
	if !utils.TokenCheck(token) {
		return errors.New("Invalid Token")
	}
	//author id 根据token得到
	authurid := repository.GetUserByToken(token) //todo 现在生成的authorid都是1
	//上传到视频服务器的视频名称
	uploadName := title
	timestamp := time.Now().Unix()
	videoName := fmt.Sprintf("%d_%s", timestamp, uploadName)
	videotype := utils.GetFormat(filename)
	videoName = videoName + "." + videotype
	fmt.Println(videoName)
	err := utils.QiniuUpload(videoName, data)
	//上传失败
	if err != nil {
		return err
	}

	playUrl := utils.TempDomainName + "/" + videoName
	//todo 视频封面选取需要改进
	coverUrl := utils.TempDomainName + "/" + "default_cover.png"
	//上传成功，写入数据库

	err = repository.PublishVideo(authurid, playUrl, coverUrl, title, timestamp)
	if err != nil {
		return err
	}
	return nil
}
func GetPublishList(token string, userid string) *PublishList {
	return &PublishList{
		Token: token,
		User_id: utils.Int64(userid, func() int64 {
			return -1
		}),
	}

}
func (v *PublishList) Do() error {
	if !utils.TokenCheck(v.Token) {
		return errors.New("输入的Token错误")
	}
	videoList, err := repository.PublishList(v.User_id)
	if err != nil {
		return err
	}
	v.VideoList = videoList
	return nil
}
