package controller

import (
	"douyinProject/entity"
	"douyinProject/repository"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type commentreturn struct {
	entity.Response
	entity.User
	Content   string `json:"content"`
	CreatData string `json:"creat_data"`
}
type comment struct {
	entity.Response
	Id          int         `json:"id"`
	User        entity.User `json:"user"`
	Content     string      `json:"content"`
	Create_data string      `json:"create_data"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	if utils.TokenCheck(token) == false { // 验证token
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  "error token",
		})
		return
	}
	id := utils.GetIdInToken(token)       // 拿到用户id
	user, _ := repository.GetUserInfo(id) // 通过用户Id拿到用户结构体
	id1, _ := strconv.Atoi(video_id)      // 获取视频id

	if action_type == "1" { // 评论操作
		com := entity.Comment{
			Id:         id,
			User:       user,
			Content:    c.Query("Content"),
			CreateDate: time.Now().String(),
		}
		var result repository.Com
		result = repository.Uploadcomment(com, id1)

		c.JSON(http.StatusOK, comment{
			Response:    entity.Response{StatusCode: 1, StatusMsg: "success"},
			Id:          int(result.Id),
			User:        user,
			Content:     result.Message,
			Create_data: result.CreateData,
		})
	} else { // 删除评论
		repository.DeleteComment(id1)
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 0,
			StatusMsg:  "delete successfully",
		})
	}
}
