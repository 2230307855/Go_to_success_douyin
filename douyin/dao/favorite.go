package dao

import (
	"douyin/models"
	"errors"

	"gorm.io/gorm"
)

// 创建一条用户点赞数据
func CreateVideoFavorite(user_id, author_id, video_id uint) error {
	// 开启事务
	tx := db.Begin()

	var favoriteRelation models.FavoriteVideoRelation
	if err := tx.Where("video_id = ? and user_id = ?", video_id, user_id).First(&favoriteRelation).Error; err == gorm.ErrRecordNotFound {
		err := tx.Create(&models.FavoriteVideoRelation{UserID: uint(user_id), VideoID: uint(video_id)}).Error
		if err != nil {
			return err
		}

		// 将 video 的 favorite_count 加 1
		if err := tx.Model(&models.Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}

		// user表中，用户喜欢总数(favorite_count)加1
		if err := tx.Model(&models.User{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}

		// user表中，视频对应的作者获赞总数(total_favorited)加1
		if err := tx.Model(&models.User{}).Where("id = ?", author_id).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
		// 保存点赞操作
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}

	return nil
}

// 取消点赞操作
// 删除用户对视频的点赞信息，被点赞视频点赞数减1
func DelFavoriteUserVideoId(userID, authorId, videoID uint) error {

	// 开启事务
	tx := db.Begin()

	// 在关联表中找到
	var FavoriteVideoRelation models.FavoriteVideoRelation
	if err := tx.Where("video_id = ? and user_id = ?", videoID, userID).First(&FavoriteVideoRelation).Error; err != nil {
		return errors.New("未找到用户和视频关系")
	} else if err == gorm.ErrRecordNotFound {
		return nil
	}

	// 删除 用户 和 视频 点赞联系
	favoriteVideoRelation := models.FavoriteVideoRelation{
		VideoID: videoID,
		UserID:  userID,
	}

	// 删除用户和被点赞视频的关系
	if err := tx.Where("video_id = ? and user_id = ?", videoID, userID).Delete(&favoriteVideoRelation).Error; err != nil {
		return err
	}

	// 视频的总点赞数减1
	if err := tx.Model(&models.Video{}).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 用户的的喜欢数量减去1
	if err := tx.Model(&models.User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// user表中，视频对应的作者获赞总数(total_favorited)减1
	if err := tx.Model(&models.User{}).Where("id = ?", authorId).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
