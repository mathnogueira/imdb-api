package movie

import "strings"

// DatabaseItem provide basic information about Movie for the database to work
type DatabaseItem struct {
	Content Movie
}

// Movie represents a movie that will be stored by this API
type Movie struct {
	ID       string
	Name     string
	Director string
	Cast     []string
}

// GetContent returns the movie it represents in the database
func (databaseItem DatabaseItem) GetContent() interface{} {
	return databaseItem.Content
}

// GetKeys generates all keys a movie can be identified by
func (databaseItem DatabaseItem) GetKeys() []string {
	keys := make([]string, 0)

	keys = append(keys, strings.Split(strings.ToLower(databaseItem.Content.Name), " ")...)
	keys = append(keys, strings.Split(strings.ToLower(databaseItem.Content.Director), " ")...)
	for _, castMember := range databaseItem.Content.Cast {
		keys = append(keys, strings.Split(strings.ToLower(castMember), " ")...)
	}

	return keys
}
