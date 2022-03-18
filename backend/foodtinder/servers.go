package foodtinder

import (
	"context"
	"io/fs"
)

// Server describes the top-level server information.
type Server interface {
	AssetServer() fs.FS
	LoginServer() LoginServer
	AuthorizerServer() AuthorizerServer
	AuthorizedServer(s *Session) AuthorizedServer
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
	// Self returns the currently authorized user.
	Self(context.Context) (*Self, error)
	// Session returns the session information.
	Session(context.Context) (*Session, error)
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
	LikedPosts(ctx context.Context) ([]Post, error)
	// DeletePosts deletes the given posts. Only the posts that belong to the
	// current user can be deleted.
	DeletePost(ctx context.Context, id ID) error
}

// UserServer describes a service serving Users.
type UserServer interface {
	// User fetches the user given the ID. Use this to fetch other users.
	User(ctx context.Context, id ID) (*User, error)
	// FoodPreferences fetches the food preferences of the user with the given
	// ID.
	FoodPreferences(ctx context.Context) (*FoodPreferences, error)
}
