package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// one record in either export file
type entry struct {
	Title          string `json:"title"`
	StringListData []struct {
		Value string `json:"value"`
	} `json:"string_list_data"`
}

// read + decode JSON file
func load(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// set of accounts you follow (titles)
func following(path string) (map[string]bool, error) {
	var f struct {
		Following []entry `json:"relationships_following"`
	}
	if err := load(path, &f); err != nil {
		return nil, err
	}
	set := make(map[string]bool, len(f.Following))
	for _, e := range f.Following {
		set[e.Title] = true
	}
	return set, nil
}

// set of accounts that follow you (across split export parts)
func followers(glob string) (map[string]bool, error) {
	files, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	set := map[string]bool{}
	for _, file := range files {
		var entries []entry
		if err := load(file, &entries); err != nil {
			return nil, err
		}
		for _, e := range entries {
			for _, s := range e.StringListData {
				set[s.Value] = true
			}
		}
	}
	return set, nil
}

func main() {
	dir := filepath.Join("connections", "followers_and_following")

	follows, err := following(filepath.Join(dir, "following.json"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fans, err := followers(filepath.Join(dir, "followers_*.json"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// following minus followers
	var notBack []string
	for name := range follows {
		if !fans[name] {
			notBack = append(notBack, name)
		}
	}
	sort.Strings(notBack)

	fmt.Printf("%d/%d don't follow back:\n", len(notBack), len(follows))
	for _, name := range notBack {
		fmt.Println(name)
	}
}
