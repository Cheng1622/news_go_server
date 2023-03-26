package models

import "time"

type Post struct {
	Status      int32  `json:"status" db:"status"`
	CommunityId int64  `json:"community_id" db:"community_id" binding:"required"`
	Id          int64  `json:"id,string" db:"post_id"`
	AuthorId    int64  `json:"author_id,string" db:"author_id"`
	Title       string `json:"title" db:"title" binding:"required" `
	Content     string `json:"content" db:"content" binding:"required" `

	Isnews     int32  `json:"isnews" db:"isnews"`
	NewsUrl    string `json:"news_url" db:"news_url"`
	NewsSource string `json:"news_source" db:"news_source"`
	NewsTime   string `json:"news_time" db:"news_time"`
	Image1     string `json:"image1" db:"image1"`
	Image2     string `json:"image2" db:"image2"`
	Image3     string `json:"image3" db:"image3"`
	Isimage    int32  `json:"isimage" db:"isimage"`
	Isimage3   int32  `json:"isimage3" db:"isimage3"`
	Videoimage string `json:"videoimage" db:"videoimage"`
	Video      string `json:"video" db:"video"`
	Isvideo    int32  `json:"isvideo" db:"isvideo"`

	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type ApiPostDetail struct {
	AuthorEmail string `json:"author_email"`
	*Community  `json:"_community"`
	*Post       `json:"_post"`
}
