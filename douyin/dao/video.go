package dao

import (
	"douyin/models"
)

// 返回所有视频
func GetAllVideos() ([]models.Video, error) {
	var videos []models.Video

	// 预加载作者信息，并查询所有视频
	result := db.Preload("Author").Limit(30).Order("created_at desc").Find(&videos)

	return videos, result.Error
}

// 发步视频-ChenglongShi
func PublishVideo(newVideo models.Video) error {
	result := db.Create(&newVideo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据用户Id查询视频关联的作者
func GetAuthorById(id int, author *models.AuthorOfVideo) error {

	result := db.Table("users").Select("id, user_name as name, following_count as follow_count, follower_count, if(follower_count>0,TRUE,FALSE) as is_follow, avatar, background_image, signature, total_favorited, work_count, favorite_count").
		Where("id = ?", id).
		First(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据用户id获取该用户的所有发布视频
func GetVideoListByAuthorId(id int, videoList *[]models.VideoItem) error {
	result := db.Table("videos").Select("id,play_url,cover_url,favorite_count,comment_count,if(favorite_count>0,TRUE,FALSE) as is_favorite,title").
		Where("author_id", id).Find(videoList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
