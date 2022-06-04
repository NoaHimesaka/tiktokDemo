package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"douyinProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfoResponse struct {
	Response entity.Response `json:"response,omitempty"`
	User     entity.User     `json:"user,omitempty"`
}

func GetUserInfoFail(err error, c *gin.Context) {
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		User: entity.User{},
	})
	log.Println(err)
}

func UserInfo(ctx *gin.Context) {
	user_Id := ctx.Query("user_id")
	fmt.Println(user_Id)
	token := ctx.Query("token")
	UserInfo, err := service.GetUserInfo(utils.Int64(user_Id, func() int64 {
		return -1
	}), token)
	if err != nil {
		GetUserInfoFail(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, UserInfoResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "成功返回数据"},
		User:     UserInfo,
	})
}
