package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()

	if err := mysql.CreatePost(p); err != nil {
		return err
	}

	return redis.CreatePost(p.ID)
}

func GetPostByID(id int64) (data *models.ApiPost, err error) {

	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID failed", zap.Error(err))
		return nil, err
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail() failed", zap.Error(err))
		return nil, err
	}

	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById() failed", zap.Error(err))
		return nil, err
	}

	data = &models.ApiPost{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(p *models.ParamPostList) ([]*models.ApiPost, error) {

	posts, err := mysql.GetPostList(p)
	if err != nil {
		return nil, err
	}

	data := make([]*models.ApiPost, 0, len(posts))

	for _, post := range posts {
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail() failed", zap.Error(err))
			continue
		}

		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById() failed", zap.Error(err))
			continue
		}

		data = append(data, &models.ApiPost{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		})
	}
	return data, err
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPost, err error) {
	ids, err := redis.GetPostIdsInOrder(p)
	if err != nil || len(ids) == 0 {
		return
	}

	posts, err := mysql.GetPostDetailByIDs(ids)
	if err != nil {
		return
	}

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	for idx, post := range posts {
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail() failed", zap.Error(err))
			continue
		}

		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById() failed", zap.Error(err))
			continue
		}

		data = append(data, &models.ApiPost{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		})
	}
	return
}
