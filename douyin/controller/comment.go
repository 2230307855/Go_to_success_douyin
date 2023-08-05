package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// author: mika

// 登录用户对视频进行评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	user_id, err1 := utils.GetIdFromToken(token)
	//验证token
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 8)
	//将评论信息上传到数据库
	if actionType == "1" {
		commentText := c.Query("comment_text")
		comment := models.Comment{
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
			VideoID:    uint(videoId),
			UserID:     uint(user_id),
			Content:    commentText,
			LikeCount:  0,
			TeaseCount: 0,
		}
		commentId := dao.CommitComment(comment)//上传评论到数据库，返回评论id
		user := models.AuthorOfVideo{}
		dao.GetAuthorById(int(comment.UserID), &user) //根据用户id查用户结构
		createDate := comment.CreatedAt.String()[5:10]
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "评论上传到数据库成功",
			"comment": models.CommentItem{
				ID:         commentId,
				User:       user,
				Content:    comment.Content,
				CreateDate: createDate,
			},
		})
	}
	//删除评论
	if actionType == "2" {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 8)
		err:=dao.DeleteComment(models.Comment{
			ID: uint(commentId),
		})
		if err!=nil{			
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusInternalServerError,
				"status_msg":  "评论删除失败",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "评论删除成功",
			})
		}

	}
}

// 查看视频的所有评论，按发布时间倒序
func CommentList(c *gin.Context) {
	token := c.Query("token")
	_, err1 := utils.GetIdFromToken(token)
	//验证token
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	videoId, _ := strconv.Atoi(c.Query("video_id"))
	var comments []models.Comment                            //用于查询的comment列表结构
	err := dao.GetCommentsByVideoId(int(videoId), &comments) //根据视频id查询评论
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "评论列表查询失败",
		})
	} else {
		var commentList []models.CommentItem //用于响应的comment列表结构
		for _, comment := range comments {
			user := models.AuthorOfVideo{}
			dao.GetAuthorById(int(comment.UserID), &user) //根据用户id查用户结构
			createDate := comment.CreatedAt.String()[5:10]
			commentList = append(commentList, models.CommentItem{
				ID:         int(comment.ID),
				User:       user,
				Content:    comment.Content,
				CreateDate: createDate,
			})
		}
		c.JSON(http.StatusOK,models.GetCommentListResponse{
			StatusCode:  0,
			StatusMsg:   "评论列表查询成功",
			CommentList: commentList,
		})
	}
}
