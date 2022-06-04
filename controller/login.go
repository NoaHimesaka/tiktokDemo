package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type loginResponse struct {
	Response entity.Response `json:"response,omitempty"`
	Id       int64           `json:"user_id,omitempty"`
	Token    string          `json:"token,omitempty"`
}

func loginFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, registerResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		Id:    -1,
		Token: "Not available",
	})
	log.Println(err)
}

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	err := service.NotContainsUsername(username)
	if err != nil {
		loginFail(ctx, err)
		return
	}
	//账户没问题，进行登录
	id, token, err := service.UserLogin(username, password)
	if err != nil {
		loginFail(ctx, err)
		return
	}

	//
	ctx.JSON(http.StatusOK, loginResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "成功返回数据"},
		Id:       id,
		Token:    token,
	})
}
