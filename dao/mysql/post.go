package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostDetailByID(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr :=
		`select post_id, title, content, author_id, community_id, create_time 
		from post 
		where post_id = ?`
	err = db.Get(data, sqlStr, id)
	return
}

func GetPostList(p *models.ParamPostList) (data []*models.Post, err error) {
	sqlStr :=
		`select post_id, title, content, author_id, community_id, create_time 
		from post
		limit ?,?`

	data = make([]*models.Post, 0, 2)
	if err := db.Select(&data, sqlStr, (p.Page-1)*p.Size, p.Size); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetPostDetailByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr :=
		`select post_id, title, content, author_id, community_id, create_time 
		from post 
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)

	return
}
