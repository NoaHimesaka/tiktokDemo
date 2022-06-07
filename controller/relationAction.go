package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"douyinProject/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(ctx *gin.Context) {
	token := ctx.Query("token")                             //用户鉴权token
	toUserId := ctx.Query("to_user_id")                     //对方用户id
	actionType, _ := strconv.Atoi(ctx.Query("action_type")) //说明：1-关注，2-取消关注
	//鉴权并拿id
	if !utils.TokenCheck(token) {
		ctx.JSON(http.StatusOK, entity.Response{
			StatusCode: -1,
			StatusMsg:  errors.New("无效Token").Error(),
		})
		return
	}
	currentUserId := utils.GetIdInToken(token)
	err := service.FollowAction(currentUserId, utils.Int64(toUserId, func() int64 {
		return -1
	}), actionType)
	if err != nil {
		ctx.JSON(http.StatusOK, entity.Response{
			StatusCode: -1,
			StatusMsg:  errors.New("crud时出现问题").Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "关注操作完成",
	})

}
