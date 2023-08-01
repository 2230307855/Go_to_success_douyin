package models

// 返回视频列表的响应模型 不是数据库中的表
// 响应体
type GetVideoListResponse struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoItem `json:"video_list"`
}

// 响应体视频列表
type VideoItem struct {
	Id            int           `json:"id"`
	Author        AuthorOfVideo `json:"author" gorm:"-"`
	PlayUrl       string        `json:"play_url"`
	CoverUrl      string        `json:"cover_url"`
	FavoriteCount int           `json:"favorites_count"`
	CommentCount  int           `json:"comment_count"`
	IsFavorite    bool          `json:"is_favorite"`
	Title         string        `json:"title"`
}

// 视频列表里的作者信息
type AuthorOfVideo struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}
