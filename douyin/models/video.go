package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time `gorm:"not null;index:idx_create" json:"created_at,omitempty"`
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	FavoritedByUsers []User         `gorm:"many2many:user_favorite_videos" json:"favorited_by_users,omitempty"`
	Author           User           `gorm:"foreignkey:AuthorID" json:"author,omitempty"`
	AuthorID         uint           `gorm:"index:idx_authorid;not null" json:"-"`
	PlayUrl          string         `gorm:"type:varchar(255);not null" json:"play_url,omitempty"`
	CoverUrl         string         `gorm:"type:varchar(255)" json:"cover_url,omitempty"`
	FavoriteCount    uint           `gorm:"default:0;not null" json:"favorite_count,omitempty"`
	CommentCount     uint           `gorm:"default:0;not null" json:"comment_count,omitempty"`
	Title            string         `gorm:"type:varchar(50);not null" json:"title,omitempty"`
	Is_favorite      bool           `gorm:"default:false" json:"is_favorite"`
}

func (Video) TableName() string {
	return "videos"
}
