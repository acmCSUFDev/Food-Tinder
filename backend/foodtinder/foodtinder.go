package foodtinder

import (
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/dataset/foods"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/time/rate"
)

// SnowflakeEpoch is the Epoch start time of a Snowflake ID used in the
// application. It is in milliseconds.
const SnowflakeEpoch = int64(1577865600 * (time.Second / time.Millisecond))

func init() { snowflake.Epoch = SnowflakeEpoch }

// AssetHash is a hash pointing to a static asset whose type is determined by
// the URL extension.
type AssetHash string

// Date describes a Date with undefined time.
type Date struct {
	D uint8
	M uint8
	Y uint16
}

// String formats Date in "02 January 2006" format.
func (d Date) String() string {
	return d.Time().Format("02 January 2006")
}

// Time returns the date in local time.
func (d Date) Time() time.Time {
	return time.Date(int(d.Y), time.Month(d.M), int(d.D), 0, 0, 0, 0, time.Local)
}

// BlurHash describes a string that is a hashed version of any image. The string
// is hashed using the Blurhash algorithm.
type BlurHash string

// ID is the Snowflake ID type for an entity. An inherent property of Snowflake
// IDs is that creation time is embedded inside the ID itself. Thus, all IDs,
// when sorted, will be sorted according to creation time. Its underlying type
// is a 64-bit signed integer.
type ID = snowflake.ID

// Throttle describes a basic rate limit throttling. It allows n bursts over a
// duration.
type Throttle struct {
	Bursts   int
	Duration time.Duration
}

// NewLimiter creates a new rate.Limiter with the Throttle description.
func (t *Throttle) NewLimiter() *rate.Limiter {
	return rate.NewLimiter(rate.Every(t.Duration), t.Bursts)
}

var (
	// UserLikedThrottle throttles the number of times the user can like a post
	// over 8 hours. It's mostly to prevent the user from spamming likes.
	UserLikedThrottle = Throttle{
		Bursts:   5,
		Duration: 8 * time.Hour,
	}
)

// LoginMetadata is the metadata of each login or register operation. As with
// all metadata, everything in this structure is optional.
type LoginMetadata struct {
	// UserAgent is the user-agent that the user was on when they logged in.
	UserAgent string
}

// Self extends User to contain personal-sensitive information.
type Session struct {
	// UserID is the ID of the user that the session identifies.
	UserID ID
	// Token is the token of the session.
	Token string
	// Expiry is the time that the session token expires.
	Expiry time.Time
	// Metadata is the metadata that was created when the user first logged in.
	Metadata LoginMetadata
}

// Self extends User and contains private information about that user.
type Self struct {
	User
	// Birthday is the user's birthday.
	Birthday Date
}

// User describes a user.
type User struct {
	ID ID
	// Name is the username which can contain spaces.
	Name string
	// Avatar is the asset hash string that can be used to create a URL.
	Avatar AssetHash
	// Bio is the user biography (or description).
	Bio string
}

// FoodPreferences describes a user's food preferences.
type FoodPreferences struct {
	// Likes is a list of categories of food that the user likes.
	Likes []foods.Category
	// Prefers maps each category that the user likes to a list of specific
	// foods that they've selected.
	Prefers map[foods.Category][]foods.Name
}

// UserLikedPosts holds the list of foods that the user liked.
type UserLikedPosts struct {
	// Posts maps post IDs to the time that the user liked.
	Posts map[ID]time.Time
	// Remaining is the number of likes allowed by the user until the Expires
	// timestamp.
	Remaining int
	// Expires is the time that the rate limiter (the Remaining field)
	// replenishes.
	Expires time.Time
}

// Post describes a posted food item.
type Post struct {
	ID ID
	// UserID is the ID of the user who posted the food item.
	UserID ID
	// CoverHash is the blur hash of the cover image.
	CoverHash BlurHash
	// Images is the list of image asset hashes for this food item. The first
	// image should be used as the cover.
	Images []AssetHash
	// Description contains the description of the food item.
	Description string
	// Tags is a list of food names that this post is relevant to. It can be
	// used for the recommendation algorithm.
	Tags []foods.Name
	// Location is the location where the post was made.
	Location string
}

// UserPostPreferences extends Post to add information specific to a single
// user.
type UserPostPreferences struct {
	Post
	// LikedAt is not nil if the user has liked the post before.
	LikedAt *time.Time
}
