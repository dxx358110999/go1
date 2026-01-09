package model

import "time"

type Article struct {
	Id       int64 `gorm:"primaryKey"` // 帖子id
	AuthorID int64 // 作者id
	//CommunityID int64     // 社区id
	//Status      int32     // 帖子状态
	Title      string    // 帖子标题
	Content    string    // 帖子内容
	CreateTime time.Time // 帖子创建时间
	UpdateTime time.Time // 帖子创建时间
}

func (Article) TableName() string {
	return "article"
}
