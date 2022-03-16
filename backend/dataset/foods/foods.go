package foods

import (
	"encoding/json"
	"log"
	"sort"
	"sync"

	_ "embed"
)

//go:generate go run ./cmd/update-food-list -o foods.json

//go:embed foods.json
var JSON []byte

var (
	once   sync.Once
	parsed Categories
)

// Unmarshal unmarshals the JSON.
func Unmarshal() Categories {
	once.Do(func() {
		if err := json.Unmarshal(JSON, &parsed); err != nil {
			log.Panicln("invalid foods.json:", err)
		}
	})
	return parsed
}

type (
	Category = string
	Name     = string

	Categories map[Category]Items
	Items      []Name
)

// CategoryNames returns the category names sorted in alphabetical order.
func (c Categories) CategoryNames() []string {
	cats := make([]string, 0, len(c))
	for cat := range c {
		if cat != "" {
			cats = append(cats, string(cat))
		}
	}
	sort.Strings(cats)
	return cats
}
