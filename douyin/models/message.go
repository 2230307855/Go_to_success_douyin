package models

import (
	"time"

	"gorm.io/gorm"
)

// 聊天信息结构
type Message struct {
	ID         uint      	  `gorm:"primarykey" json:"id"`
	ToUser     User           `gorm:"foreignkey:ToUserID;" json:"-"`
	ToUserID   uint           `gorm:"index:idx_userid_to;not null" json:"to_user_id"`
	FromUser   User           `gorm:"foreignkey:FromUserID;" json:"-"`
	FromUserID uint           `gorm:"index:idx_userid_from;not null" json:"from_user_id"`
	Content    string         `gorm:"type:varchar(255);not null" json:"content"`
	CreatedAt  time.Time 	  `gorm:"index:created_at;not null" json:"create_time"`
	UpdatedAt  time.Time	  `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Message) TableName() string {
	return "messages"
}
