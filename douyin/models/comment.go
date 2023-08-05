package models

import (
	"time"
	"gorm.io/gorm"
)
// author: mika

// 用户评论数据模型，用于数据库
type Comment struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `gorm:"index;not null" json:"create_date"`
	UpdatedAt  time.Time 
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Video      Video          `gorm:"foreignkey:VideoID" json:"video,omitempty"`
	VideoID    uint           `gorm:"index:idx_videoid;not null" json:"video_id"`
	User       User           `gorm:"foreignkey:UserID" json:"user,omitempty"`
	UserID     uint           `gorm:"index:idx_userid;not null" json:"user_id"`
	Content    string         `gorm:"type:varchar(255);not null" json:"content"`
	LikeCount  uint           `gorm:"column:like_count;default:0;not null" json:"like_count,omitempty"`
	TeaseCount uint           `gorm:"column:tease_count;default:0;not null" json:"tease_count,omitempty"`
}

//查询评论列表的响应包
type GetCommentListResponse struct {
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
	CommentList []CommentItem `json:"comment_list"`
}

//用于响应的Comment结构，使用AuthorOfVideo作为用户模板
type CommentItem struct {
	ID int `json:"id"`
	User AuthorOfVideo `json:"user"`
	Content string `json:"content"`
	CreateDate string `json:"create_date"`
}
