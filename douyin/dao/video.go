package dao

import "douyin/models"

// 返回所有视频
func GetAllVideos() ([]models.Video, error) {
	var videos []models.Video

	// 预加载作者信息，并查询所有视频
	result := db.Preload("FavoritedByUsers").Preload("Author").Limit(30).Order("created_at desc").Find(&videos)

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

// 根据视频id，返回作者id
func VideoAuthorId(video_id uint) (uint, error) {
	var videos models.Video

	// 预加载作者信息，并查询作者ID
	if err := db.Preload("Author").First(&videos, video_id).Error; err != nil {
		return 0, err
	}

	return videos.AuthorID, nil

}
