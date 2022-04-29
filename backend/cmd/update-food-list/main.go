package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/dataset/foods"
)

const src = "https://raw.githubusercontent.com/Open-Food/Open-Food-Standard/master/datasets/Food%20Types%207%20Feb%202012.csv"

var (
	output = "-"
)

func init() {
	flag.StringVar(&output, "o", output, "output path, - for stdout")
	flag.Parse()

	http.DefaultClient.Timeout = 10 * time.Second
}

func main() {
	out := os.Stdout

	if output != "-" {
		s, err := os.Stat(output)
		if err == nil && !s.IsDir() {
			log.Printf("%s already exists, skipping", output)
			return
		}

		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Fatalln("cannot make output file:", err)
		}
		defer f.Close()

		out = f
	}

	r, err := http.Get(src)
	if err != nil {
		log.Fatalln("cannot get food_types.rtf:", err)
	}
	defer r.Body.Close()

	strmap, err := scanFoods(r.Body)
	if err != nil {
		log.Fatalln("cannot scan food_types.rtf:", err)
	}

	enc := json.NewEncoder(out)
	enc.SetIndent("", " ")

	if err := enc.Encode(strmap); err != nil {
		log.Fatalln("cannot encode stringMap:", err)
	}
}

type (
	foodCategories map[string]foodSet
	foodSet        map[string]struct{}
)

func (categories foodCategories) add(category string, items ...string) {
	list, ok := categories[category]
	if !ok {
		list = make(foodSet, 20)
		categories[category] = list
	}

	for _, item := range items {
		if item == "" {
			continue
		}
		item = sanitizeName(item)
		list[item] = struct{}{}
	}
}

func (categories foodCategories) toList() foods.Categories {
	m := make(foods.Categories, len(categories))
	for k, set := range categories {
		m[foods.Category(k)] = set.toList()
	}
	return m
}

func (set foodSet) toList() foods.Items {
	items := make(foods.Items, 0, len(set))
	for item := range set {
		items = append(items, foods.Name(item))
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	return items
}

var lowerFirstLetterRe = regexp.MustCompile(`(^|[ (/])[a-z]`)

func sanitizeName(name string) string {
	if name == "" {
		return name
	}

	// This messes up certain stuff.
	// item = strings.Title(item)
	name = strings.TrimSpace(name)
	// That one edge case.
	name = strings.ReplaceAll(name, "�", "'")
	// Sanitize casing.
	name = lowerFirstLetterRe.ReplaceAllStringFunc(name, strings.ToUpper)
	// No idea why this is here.
	name = strings.TrimSuffix(name, "-")

	// Reverse this like "last name, first name".
	if strings.Count(name, ", ") == 1 {
		parts := strings.Split(name, ", ")
		name = parts[1] + " " + parts[0]
	}

	return name
}

func scanFoods(src io.Reader) (foods.Categories, error) {
	r := csv.NewReader(src)
	r.TrimLeadingSpace = true

	_, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("cannot read CSV header: %q", err)
	}

	if r.FieldsPerRecord < 7 {
		return nil, fmt.Errorf("missing columns, have %d, want %d", r.FieldsPerRecord, 7)
	}

	categories := make(foodCategories, 25)

	for {
		cols, err := r.Read()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
			break
		}

		// Use the name (0) and the simple name (5).
		names := []string{cols[0], cols[5]}
		// 3rd column (2) is a comma-delimited list (inside a csv file!) of
		// alternative names. Use that if available.
		for _, name := range strings.Split(cols[2], ",") {
			names = append(names, strings.TrimSpace(name))
		}

		// Sanitize the inconsistent casing.
		category := strings.Title(cols[1])
		categories.add(category, names...)
	}

	return categories.toList(), nil
}
