package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response  entity.Response `json:"response,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
	VideoList []entity.Video  `json:"video_list,omitempty"`
}

func Feedfail(err error, c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		NextTime:  time.Now().Unix(),
		VideoList: nil,
	})
	log.Println(err)
}

func Feed(ctx *gin.Context) {
	latestTime := ctx.Query("latest_time") //限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	token := ctx.Query("token")            //用户登录状态下设置
	VideoInstances := service.NewVideoInstance(latestTime, token)
	err := VideoInstances.Do()
	if err != nil {
		Feedfail(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, FeedResponse{
		Response:  entity.Response{StatusCode: 0, StatusMsg: "成功返回数据"},
		NextTime:  time.Now().Unix(),
		VideoList: *(VideoInstances.VideoList),
	})
}
