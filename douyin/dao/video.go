package dao

import "douyin/models"

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
