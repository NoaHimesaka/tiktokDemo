package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type publishlistResponse struct {
	Response  entity.Response `json:"response,omitempty"`
	VideoList []entity.Video  `json:"video_list,omitempty"`
}

func Publishlistfail(err error, c *gin.Context) {
	c.JSON(http.StatusOK, publishlistResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
	})
	log.Println(err)
}

func Publishlist(ctx *gin.Context) {
	userid := ctx.Query("user_id")
	token := ctx.Query("token")
	publishlistInstance := service.GetPublishList(token, userid)
	err := publishlistInstance.Do()
	if err != nil {
		Publishlistfail(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, publishlistResponse{
		Response:  entity.Response{StatusCode: 0, StatusMsg: "成功返回数据"},
		VideoList: *(publishlistInstance.VideoList),
	})
}
