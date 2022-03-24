package foodtinder

import (
	"context"
	"io"
	"io/fs"
)

// Server describes the top-level server information.
type Server interface {
	FileServer() FileServer
	LoginServer() LoginServer
	AuthorizerServer() AuthorizerServer
	AuthorizedServer(s *Session) AuthorizedServer
}

// FileServer describes a service for serving and creating files.
//
// TODO: consider if we should keep track of which user uploaded which asset.
// Otherwise, we can smartly reference-count them.
type FileServer interface {
	fs.FS
	// fs.StatFS

	// Create creates a file with the content from the given Reader. The new
	// asset hash is returned, or an error if there's one. Create may be called
	// concurrently; the implementation should guarantee that everything is
	// atomic and that there is no collision.
	Create(s *Session, src io.Reader) (string, error)
}

// LoginServer describes a service serving Sessions.
type LoginServer interface {
	// Login authenticates the user by a username and password and returns a new
	// session.
	Login(ctx context.Context, username, password string, m LoginMetadata) (*Session, error)
	// Register registers a new user.
	Register(ctx context.Context, username, password string, m LoginMetadata) (*Session, error)
}

// AuthorizerServer describes a service for authorizing a session.
type AuthorizerServer interface {
	// Authorize authorizes the user using the given session token. The session
	// is returned if the token points to a valid user.
	Authorize(ctx context.Context, token string) (*Session, error)
}

// AuthorizedServer describes a service for a specific user session.
type AuthorizedServer interface {
	// Logout invalidates the authorizing token.
	Logout(context.Context) error

	PostServer() PostServer
	UserServer() UserServer
}

// PostServer is a service serving Posts.
type PostServer interface {
	// NextPosts returns the list of posts that the user will see next
	// starting from the given previous ID. If the ID is 0, then the top is
	// assumed.
	NextPosts(ctx context.Context, previousID ID) ([]Post, error)
	// LikedPosts returns the list of liked posts by the user.
	LikedPosts(ctx context.Context) ([]Post, error)
	// DeletePosts deletes the given posts. Only the posts that belong to the
	// current user can be deleted.
	DeletePost(ctx context.Context, id ID) error
	// CreatePost creates a new post. The post's ID is ignored and a new one is
	// used. That ID is returned back.
	CreatePost(ctx context.Context, post Post) (ID, error)
}

// UserServer describes a service serving Users.
type UserServer interface {
	// User fetches the user given the ID. Use this to fetch other users.
	User(ctx context.Context, username string) (*User, error)
	// Self returns the current user.
	Self(context.Context) (*Self, error)
	// UpdateSelf updates the current user. This method should change everything
	// except the Username and Password.
	UpdateSelf(context.Context, *Self) error
	// ChangePassword changes the password of the current user. All existing
	// sessions except for this one should be invalidated.
	ChangePassword(ctx context.Context, newPassword string) error
}

// MaxAssetSize is the maximum size in bytes that an asset should be.
const MaxAssetSize = 1 << 20 // 1MB
