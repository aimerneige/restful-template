package models

type ArticleDto struct {
	ID       uint   `json:"id"`
	CreateAt string `json:"create"`
	UpdateAt string `json:"update"`
	UserID   uint   `json:"userid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Likes    uint   `json:"likes"`
}
