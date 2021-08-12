package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserID  uint
	Title   string
	Content string
	Likes   uint
}

func (a Article) ToDto() (dto ArticleDto) {
	timeTemplate := "2006-01-02 15:04"

	dto.ID = a.ID
	dto.CreateAt = a.CreatedAt.Local().Format(timeTemplate)
	dto.UpdateAt = a.UpdatedAt.Local().Format(timeTemplate)
	dto.UserID = a.UserID
	dto.Title = a.Title
	dto.Content = a.Content
	dto.Likes = a.Likes

	return
}
