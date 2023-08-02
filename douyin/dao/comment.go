package dao

import (
	"douyin/models"
	"sort"
)

// author: mika

//向数据库增加一条comment
func CommitComment(comment models.Comment)int{
	db.Create(&comment)
	return int(comment.ID)
}
//删除一条comment，gorm使用软删除，只是做标记
func DeleteComment(comment models.Comment)error{
	result:=db.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
//根据视频id查询发表的评论，按发布时间倒序
func GetCommentsByVideoId(videoId int, comments *[]models.Comment)error{
	result := db.Table("comments").Select("id,user_id,content,created_at").Where("video_id", videoId).Find(comments)
	sort.Slice(*comments,func(i,j int) bool{
		return (*comments)[j].CreatedAt.Before((*comments)[i].CreatedAt)
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}