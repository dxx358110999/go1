package dto

type Article struct {
	Id       int64 // 帖子id
	AuthorID int64 `json:"author_id" binding:"required"` // 作者id
	//CommunityID int64     // 社区id
	//Status      int32     // 帖子状态
	Title   string `json:"title" binding:"required"`   // 帖子标题
	Content string `json:"content" binding:"required"` // 帖子内容
}
