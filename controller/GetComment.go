package controller

import (
	"douyinProject/entity"
	"douyinProject/repository"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type comm struct {
	entity.Response
	List []repository.Result `json:"comment_list"`
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	if utils.TokenCheck(token) == true {
		id, _ := strconv.Atoi(videoId) // 获取视频id
		var res comm
		res.List = repository.Get(id) // 获取评论
		j := 0
		for i := len(res.List) - 1; i >= 0 && j <= i; i-- { // 排序
			res.List[i], res.List[j] = res.List[j], res.List[i]
			j++
		}
		c.JSON(http.StatusOK, comm{
			Response: entity.Response{StatusCode: 0, StatusMsg: "ok"},
			List:     res.List,
		})
	} else {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  "error token",
		})
	}
}
