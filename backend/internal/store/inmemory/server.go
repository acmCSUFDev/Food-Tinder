// Package memory implements foodtinder.Server using on-memory maps.
package inmemory

import (
	"io/fs"
	"log"
	"sort"
	"sync"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/bwmarrin/snowflake"
)

// NewID creates a new ID with the given order.
func NewID(order int) foodtinder.ID {
	n, err := snowflake.NewNode(int64(order))
	if err != nil {
		log.Panicln("NewID error:", err)
	}
	return n.Generate()
}

// User describes an on-memory user.
type User struct {
	foodtinder.Self
	Password    string
	Preferences foodtinder.FoodPreferences
	LikedPosts  []foodtinder.ID
}

type Session struct {
	Username string
	Metadata foodtinder.LoginMetadata
}

// State describes the initial state of the memory store.
type State struct {
	Users    []User
	Posts    []foodtinder.Post
	Sessions []foodtinder.Session
	Assets   map[string][]byte
}

func (s *State) sortPosts() {
	sort.Slice(s.Posts, func(i, j int) bool {
		return s.Posts[i].ID < s.Posts[j].ID
	})
}

type server struct {
	mu       sync.RWMutex
	store    State
	sessions map[string]*foodtinder.Session
	assetIx  uint64
}

// NewServer creates a new on-memory data server using the given store.
func NewServer(state State) foodtinder.Server {
	state.sortPosts()

	if state.Assets == nil {
		state.Assets = make(map[string][]byte)
	}

	sessions := make(map[string]*foodtinder.Session)
	for i, session := range state.Sessions {
		sessions[session.Token] = &state.Sessions[i]
	}

	return &server{
		store:    state,
		sessions: sessions,
	}
}

func (s *server) AssetServer() fs.FS {
	return (*assetServer)(s)
}

func (s *server) LoginServer() foodtinder.LoginServer {
	return (*loginServer)(s)
}

func (s *server) AuthorizerServer() foodtinder.AuthorizerServer {
	return (*authorizerServer)(s)
}

func (s *server) AuthorizedServer(session *foodtinder.Session) foodtinder.AuthorizedServer {
	return &authorizedServer{
		server:  s,
		session: session,
	}
}
