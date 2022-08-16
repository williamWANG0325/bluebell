package models

import "time"

type Post struct {
	ID          int64     `db:"post_id" json:"id,string"`
	AuthorID    int64     `db:"author_id" json:"author_id,string"`
	CommunityID int64     `db:"community_id" json:"community_id" binding:"required"`
	Status      int32     `db:"status" json:"status"`
	Title       string    `db:"title" json:"title" binding:"required"`
	Content     string    `db:"content" json:"content" binding:"required"`
	CreateTime  time.Time `db:"create_time" json:"create_time"`
}

type ApiPost struct {
	AuthorName       string `json:"author_name"`
	VoteNum          int64  `json:"vote_num"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community_detail"`
}
