package routes

import (
	"douyin/controller"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {

	apiRouter := r.Group("douyin")

	// 测试
	apiRouter.GET("/test/", controller.Test)

	// Basic
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)            // 用户信息
	apiRouter.POST("/user/register/", controller.Register)  // 注册
	apiRouter.POST("/user/login/", controller.Login)        // 登录
	apiRouter.POST("/publish/action/", controller.Publish)  // 登录用户选择视频上传
	apiRouter.GET("/publish/list/", controller.PublishList) // 用户的视频发布列表，直接列出用户所有投稿过的视频

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction) // 登录用户对视频的点赞和取消点赞操作
	apiRouter.GET("/favorite/list/", controller.FavoriteList)      // 用户的所有点赞视频
	apiRouter.POST("/comment/action/", controller.CommentAction)   // 登录用户对视频进行评论
	apiRouter.GET("/comment/list/", controller.CommentList)        // 查看视频的所有评论，按发布时间倒序

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)     // 关注操作
	apiRouter.GET("/relation/follow/list/", controller.FollowList)     // 关注列表
	apiRouter.GET("/relation/follower/list/", controller.FollowerList) // 粉丝列表
	apiRouter.GET("/relation/friend/list/", controller.FriendList)     // 好友列表
	apiRouter.GET("/message/chat/", controller.MessageChat)            // 发送信息
	apiRouter.POST("/message/action/", controller.MessageAction)       // 聊天记录
}
