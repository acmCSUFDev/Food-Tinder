// Package memory implements foodtinder.Server using on-memory maps.
package inmemory

import (
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/fileserver"
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
	Users     []User
	Posts     []foodtinder.Post
	Sessions  []foodtinder.Session
	AssetURLs []string
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
	fserver  foodtinder.FileServer
}

// NewServer creates a new on-memory data server using the given store.
func NewServer(state State, fserver foodtinder.FileServer) foodtinder.Server {
	if fserver == nil {
		fserver = fileserver.InMemory(nil)
	}

	state.sortPosts()

	sessions := make(map[string]*foodtinder.Session)
	for i, session := range state.Sessions {
		sessions[session.Token] = &state.Sessions[i]
	}

	if len(state.AssetURLs) > 0 {
		for _, url := range state.AssetURLs {
			go func(url string) {
				r, err := http.Get(url)
				if err != nil {
					log.Println("cannot GET asset", url)
					return
				}

				defer r.Body.Close()

				v, err := fserver.Create(nil, r.Body)
				if err != nil {
					log.Println("cannot download asset", url)
					return
				}

				log.Println("registered asset", v)
			}(url)
		}
	}

	return &server{
		store:    state,
		sessions: sessions,
		fserver:  fserver,
	}
}

func (s *server) FileServer() foodtinder.FileServer {
	return s.fserver
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
