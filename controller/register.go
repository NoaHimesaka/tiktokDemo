package controller

import (
	"douyinProject/entity"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type registerResponse struct {
	Response entity.Response `json:"response,omitempty"`
	Id       int64           `json:"user_id,omitempty"`
	Token    string          `json:"token,omitempty"`
}

func RegisterFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, registerResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		Id:    -1,
		Token: "",
	})
	log.Println(err)
}

func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	// todo 判断账户是否存在,如果存在就返回err
	err := service.ContainsUsername(username)
	if err != nil {
		RegisterFail(ctx, err)
		return
	}
	//没问题，进行注册
	id, token, err := service.UserRegister(username, password)
	if err != nil {
		RegisterFail(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, registerResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "成功返回数据"},
		Id:       id,
		Token:    token,
	})
}
