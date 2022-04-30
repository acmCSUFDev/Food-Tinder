// Package memory implements foodtinder.Server using on-memory maps.
package inmemory

import (
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
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

func (u *User) likes(postID foodtinder.ID) bool {
	for _, id := range u.LikedPosts {
		if id == postID {
			return true
		}
	}
	return false
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
func NewServer(state State) foodtinder.Server {
	state.sortPosts()

	sessions := make(map[string]*foodtinder.Session)
	for i, session := range state.Sessions {
		sessions[session.Token] = &state.Sessions[i]
	}

	return &server{
		store:    state,
		sessions: sessions,
		fserver:  fileserver.InMemory(fetchAssetURLs(state.AssetURLs)),
	}
}

func fetchAssetURLs(urls []string) map[string][]byte {
	if len(urls) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	assets := make(map[string][]byte, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()

			r, err := http.Get(url)
			if err != nil {
				log.Println("cannot GET asset", url)
				return
			}

			defer r.Body.Close()

			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("cannot download asset", url)
				return
			}

			mu.Lock()
			assets[strconv.Itoa(i+1)] = b
			mu.Unlock()

			log.Println("registered asset", i)
		}(i, url)
	}

	wg.Wait()
	return assets
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
