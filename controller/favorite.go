package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var msg = [2]string{"点赞成功", "取消点赞"}

type FavoriteListResponce struct {
	Response  entity.Response `json:"response,omitempty"`
	VideoList []entity.Video  `json:"video_list"`
}

func FavoriteAction(ctx *gin.Context) {
	video_id := ctx.Query("video_id")
	userId := ctx.Query("user_id")
	token := ctx.Query("token")
	action_type := ctx.Query("action_type")
	act, err := service.FavoriteAction(userId, token, video_id, action_type)
	if err != nil {
		ctx.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  msg[act-1],
	})
}
func FavoriteList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	token := ctx.Query("token")
	video, err := service.GetFavorite(userId, token)
	if err != nil {
		ctx.JSON(http.StatusOK, FavoriteListResponce{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			VideoList: []entity.Video{},
		})
	}
	ctx.JSON(http.StatusOK, FavoriteListResponce{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Success",
		},
		VideoList: *video,
	})
}
