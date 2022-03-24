package postgres

import (
	"context"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres/query"
	"github.com/pkg/errors"
)

type PostServer AuthorizedServer

func (s *PostServer) NextPosts(ctx context.Context, previousID foodtinder.ID) ([]foodtinder.Post, error) {
	r, err := s.queries.NextPosts(ctx, int64(previousID))
	if err != nil {
		return nil, errors.Wrap(err, "SQL error")
	}

	posts := make([]foodtinder.Post, len(r))
	for i, p := range r {
		posts[i] = foodtinder.Post{
			ID:          foodtinder.ID(p.ID),
			Username:    p.Username,
			CoverHash:   p.CoverHash.String,
			Images:      p.Images,
			Description: p.Description.String,
			Tags:        p.Tags,
			Location:    p.Location.String,
			Likes:       int(p.Likes),
		}
	}

	return posts, nil
}

func (s *PostServer) LikedPosts(ctx context.Context) ([]foodtinder.Post, error) {
	r, err := s.queries.LikedPosts(ctx, s.session.Username)
	if err != nil {
		return nil, errors.Wrap(err, "SQL error")
	}

	posts := make([]foodtinder.Post, len(r))
	for i, p := range r {
		posts[i] = foodtinder.Post{
			ID:          foodtinder.ID(p.ID),
			Username:    p.Username,
			CoverHash:   p.CoverHash.String,
			Images:      p.Images,
			Description: p.Description.String,
			Tags:        p.Tags,
			Location:    p.Location.String,
			Likes:       int(p.Likes),
		}
	}

	return posts, nil
}

func (s *PostServer) DeletePost(ctx context.Context, id foodtinder.ID) error {
	n, err := s.queries.DeletePost(ctx, query.DeletePostParams{
		ID:       int64(id),
		Username: s.session.Username,
	})
	if err != nil {
		return errors.Wrap(err, "SQL error")
	}
	if n == 0 {
		return foodtinder.ErrNotFound
	}
	return nil
}

func (s *PostServer) CreatePost(ctx context.Context, post foodtinder.Post) (foodtinder.ID, error) {
	id := s.idNode.Generate()

	err := s.queries.CreatePost(ctx, query.CreatePostParams{
		ID:          int64(id),
		Username:    post.Username,
		CoverHash:   nullstr(post.CoverHash),
		Images:      post.Images,
		Description: nullstr(post.Description),
		Tags:        post.Tags,
		Location:    nullstr(post.Location),
	})
	if err != nil {
		return 0, errors.Wrap(err, "SQL error")
	}

	return id, nil
}
