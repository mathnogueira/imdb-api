package database

// Item represents any item that can be stored in the storage
type Item interface {
	GetKeys() []string
	GetContent() interface{}
}

// Storage is where all items will be stored. This provides a generic interface
// for storing, retrieving, and search items
type Storage interface {
	Save(item Item) error
	Get(key string) []*Item
	Search(keys []string) []Item
}
