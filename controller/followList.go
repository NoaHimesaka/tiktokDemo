package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"douyinProject/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowListResponse struct {
	Response entity.Response `json:"response,omitempty"`
	UserList *[]entity.User  `json:"user_list,omitempty"`
}

func FollowList(ctx *gin.Context) {
	token := ctx.Query("token")    //用户鉴权token
	userId := ctx.Query("user_id") //要查看关注列表的
	//鉴权并拿id
	if !utils.TokenCheck(token) {
		ctx.JSON(http.StatusOK, entity.Response{
			StatusCode: -1,
			StatusMsg:  errors.New("无效Token").Error(),
		})
		return
	}
	//TODO:如果当前用户ID需要使用用来判断是否有权限看对方的信息的时候再解开
	//currentUserId := utils.GetIdInToken(token)
	followList, err := service.FollowListAction(utils.Int64(userId, func() int64 {
		return -1
	}))
	if err != nil {
		ctx.JSON(http.StatusOK, entity.Response{
			StatusCode: -1,
			StatusMsg:  errors.New("获取列表出现问题").Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, FollowListResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "获取列表成功"},
		UserList: &followList,
	})
}
