package inmemory

import (
	"context"
	"fmt"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/bwmarrin/snowflake"
)

type postServer authorizedServer

const pageSize = 10

func (s *postServer) Post(ctx context.Context, id foodtinder.ID) (*foodtinder.PostListing, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, err := (*userServer)(s).self()
	if err != nil {
		return nil, err
	}

	p, ok := s.post(id)
	if ok {
		p := *p
		p.Likes = s.countLikes(id)

		return &foodtinder.PostListing{
			Post:  p,
			Liked: u.likes(p.ID),
		}, nil
	}

	return nil, foodtinder.ErrNotFound
}

func (s *postServer) NextPosts(ctx context.Context, prevID foodtinder.ID) ([]foodtinder.PostListing, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, err := (*userServer)(s).self()
	if err != nil {
		return nil, err
	}

	var startIx int

	if prevID != 0 {
		for i, post := range s.store.Posts {
			if post.ID == prevID {
				startIx = i
				goto ret
			}
		}
		return nil, fmt.Errorf("unknown post with id %v", prevID)
	}

ret:
	endIx := startIx + pageSize
	if endIx > len(s.store.Posts) {
		endIx = len(s.store.Posts)
	}

	posts := s.store.Posts[startIx:endIx]
	listings := make([]foodtinder.PostListing, len(posts))

	for i, post := range posts {
		post.Likes = s.countLikes(post.ID)

		listings[i] = foodtinder.PostListing{
			Post:  post,
			Liked: u.likes(post.ID),
		}
	}

	return listings, nil
}

func (s *postServer) LikedPosts(ctx context.Context) ([]foodtinder.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, err := (*userServer)(s).self()
	if err != nil {
		return nil, err
	}

	posts := make([]foodtinder.Post, 0, len(u.LikedPosts))
	for _, id := range u.LikedPosts {
		// Ideally, this ok check should be an ON DELETE CASCADE.
		p, ok := s.post(id)
		if ok {
			posts = append(posts, *p)
		}
	}

	return posts, nil
}

func (s *postServer) LikePost(ctx context.Context, id foodtinder.ID, like bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, err := (*userServer)(s).self()
	if err != nil {
		return err
	}

	if _, ok := s.post(id); !ok {
		return foodtinder.ErrNotFound
	}

	if like {
		if !u.likes(id) {
			u.LikedPosts = append(u.LikedPosts, id)
		}
		return nil
	}

	for i, likedID := range u.LikedPosts {
		if likedID == id {
			u.LikedPosts = append(u.LikedPosts[:i], u.LikedPosts[i+1:]...)
			return nil
		}
	}

	return nil
}

func (s *postServer) post(id foodtinder.ID) (*foodtinder.Post, bool) {
	for i, post := range s.store.Posts {
		if post.ID == id {
			return &s.store.Posts[i], true
		}
	}
	return nil, false
}

func (s *postServer) countLikes(postID foodtinder.ID) int {
	var likes int
userLoop:
	for _, user := range s.store.Users {
		for _, id := range user.LikedPosts {
			if id == postID {
				likes++
				continue userLoop
			}
		}
	}
	return likes
}

func (s *postServer) DeletePost(ctx context.Context, id foodtinder.ID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, post := range s.store.Posts {
		if post.ID == id {
			s.store.Posts = append(s.store.Posts[:i], s.store.Posts[i+1:]...)
			return nil
		}
	}
	return foodtinder.ErrNotFound
}

var postIDNode *snowflake.Node

func init() {
	n, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}
	postIDNode = n
}

func (s *postServer) CreatePost(ctx context.Context, post foodtinder.Post) (foodtinder.ID, error) {
	// TODO: this is prone to the Two Generals' Problem, but it shouldn't be an
	// issue in the SQL database, at least minimally. It may still be an isuse
	// on response write, but whatever. We might want to look into
	// Idempotency-Key.

	id := postIDNode.Generate()

	s.mu.Lock()
	defer s.mu.Unlock()

	u, err := (*userServer)(s).self()
	if err != nil {
		return 0, err
	}

	post.ID = foodtinder.ID(id)
	post.Username = u.Username

	s.store.Posts = append(s.store.Posts, post)
	return post.ID, nil
}
