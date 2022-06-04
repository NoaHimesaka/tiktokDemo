package controller

import (
	"bytes"
	"douyinProject/entity"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func PublishFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK,
		entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error()},
	)
	log.Println(err)
}

func Publish(ctx *gin.Context) {
	//get token,title,data
	token := ctx.PostForm("token")
	title := ctx.PostForm("title")
	file, header, err := ctx.Request.FormFile("data")
	defer file.Close()
	if err != nil {
		PublishFail(ctx, err)
		return
	}
	filename := header.Filename
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		PublishFail(ctx, err)
		return
	}
	err = service.VideoPublish(token, title, buf.Bytes(), filename)
	if err != nil {
		PublishFail(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK,
		entity.Response{StatusCode: 0, StatusMsg: "成功上传视频"},
	)
}
