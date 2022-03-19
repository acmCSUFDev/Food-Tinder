package foodtinder

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/acmCSUFDev/Food-Tinder/backend/dataset/foods"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/runeutil"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/time/rate"
)

// SnowflakeEpoch is the Epoch start time of a Snowflake ID used in the
// application. It is in milliseconds.
const SnowflakeEpoch = int64(1577865600 * (time.Second / time.Millisecond))

func init() { snowflake.Epoch = SnowflakeEpoch }

// Error constants.
var (
	// ErrNotFound is used when a resource is not found.
	ErrNotFound = errors.New("not found")
	// ErrInvalidLogin is returned when a user logs in with an unknown
	// combination of username and password.
	ErrInvalidLogin = errors.New("invalid username or password")
	// ErrUsernameExists is returned by Register if the user tries to make an
	// account with an existing username.
	ErrUsernameExists = errors.New("an account with the username already exists")
)

// Date describes a Date with undefined time.
type Date struct {
	D uint8
	M uint8
	Y uint16
}

// IsZero returns true if Date is zero-value (is 00/00/0000).
func (d Date) IsZero() bool { return d == (Date{}) }

// String formats Date in "02 January 2006" format.
func (d Date) String() string {
	return d.Time().Format("02 January 2006")
}

// Time returns the date in local time.
func (d Date) Time() time.Time {
	return time.Date(int(d.Y), time.Month(d.M), int(d.D), 0, 0, 0, 0, time.Local)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d == (Date{}) {
		return []byte("null"), nil
	}
	return json.Marshal(fmt.Sprintf("%04d/%02d/%02d", d.Y, d.M, d.D))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*d = Date{}
		return nil
	}

	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	t, err := time.Parse("2006/01/02", str)
	if err != nil {
		return fmt.Errorf("invalid time: %w", err)
	}

	*d = Date{
		D: uint8(t.Day()),
		M: uint8(t.Month()),
		Y: uint16(t.Year()),
	}
	return nil
}

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
	// Username identifies the user that the session belongs to.
	Username string
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

// AllowedUsernameRunes is a list of validators that validate a username. Its
// rules can be described as "letters and digits only with ., _, -, *, $ and !".
var AllowedUsernameRunes = []runeutil.Validator{
	unicode.IsLetter,
	unicode.IsDigit,
	runeutil.AllowRunes('_', '-', '*', '$', '!', '.'),
}

// ValidateUsername validates the username. Usernames must be 35 characters long
// and only contain runes that satisfy AllowedUsernameRunes.
func ValidateUsername(username string) error {
	if len(username) > 35 {
		return fmt.Errorf("username too long, max 35 characters")
	}
	if runeutil.ContainsIllegal(username, AllowedUsernameRunes) {
		return fmt.Errorf("username contains illegal characters")
	}
	return nil
}

// User describes a user.
type User struct {
	// Username is the username which can contain spaces. All usernames must be
	// unique.
	Username string
	// DisplayName is the visible name that users actually see.
	DisplayName string
	// Avatar is the asset hash string that can be used to create a URL.
	Avatar string
	// Bio is the user biography (or description).
	Bio string
}

// Validate validates User.
func (u User) Validate() error {
	if len(u.Bio) > 4096 {
		return fmt.Errorf("bio too long, maximum 4096 bytes")
	}

	if err := ValidateUsername(u.Username); err != nil {
		return err
	}

	if utf8.RuneCountInString(u.DisplayName) > 50 {
		return fmt.Errorf("display name too long, must have 50 or fewer characters")
	}

	return nil
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
	// Username is the username of the user who posted the food item.
	Username string
	// CoverHash is the blur hash of the cover image.
	CoverHash string
	// Images is the list of image asset hashes for this food item. The first
	// image should be used as the cover.
	Images []string
	// Description contains the description of the food item.
	Description string
	// Tags is a list of food names that this post is relevant to. It can be
	// used for the recommendation algorithm.
	Tags []foods.Name
	// Location is the location where the post was made.
	Location string
}

// Validate validates the Post.
func (p Post) Validate() error {
	if len(p.Description) > 4096 {
		return fmt.Errorf("description too long, max 4096 characters")
	}

	if len(p.Images) > 12 {
		return fmt.Errorf("too many images, max 12 images")
	}

	return nil
}

// UserPostPreferences extends Post to add information specific to a single
// user.
type UserPostPreferences struct {
	Post
	// LikedAt is not nil if the user has liked the post before.
	LikedAt *time.Time
}
