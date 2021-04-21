package database

import (
	"strings"
	"sync"

	"go.uber.org/zap"
)

// MemoryStorage is an implementation of the interface Storage that doesn't persist any data
// on secondary memory. It relies on primary memory only.
type MemoryStorage struct {
	buckets map[string][]*Item
	logger  *zap.Logger
}

// NewMemoryStorage creates a new in-memory storage
func NewMemoryStorage(logger *zap.Logger) Storage {
	return &MemoryStorage{
		logger:  logger,
		buckets: make(map[string][]*Item, 0),
	}
}

// Save saves an item
func (memoryStorage *MemoryStorage) Save(item Item) error {
	keys := getLowercaseKeys(item.GetKeys())
	memoryStorage.logger.Debug("Saving new item in storage",
		zap.Strings("keys", keys),
	)
	for _, key := range keys {
		if bucket, bucketExists := memoryStorage.buckets[key]; bucketExists {
			if !doesItemAlreadyExistsOnBucket(bucket, &item) {
				memoryStorage.buckets[key] = append(memoryStorage.buckets[key], &item)
			}
		} else {
			memoryStorage.buckets[key] = []*Item{&item}
		}
	}

	return nil
}

func doesItemAlreadyExistsOnBucket(bucket []*Item, item *Item) bool {
	for _, bucketItem := range bucket {
		if item == bucketItem {
			return true
		}
	}

	return false
}

// Get retrieves all items that are identified by the key
func (memoryStorage *MemoryStorage) Get(key string) []*Item {
	key = strings.ToLower(key)
	items := make([]*Item, 0)
	if bucket, bucketExists := memoryStorage.buckets[key]; bucketExists {
		for _, item := range bucket {
			items = append(items, item)
		}
	}

	return items
}

// Search returns all items that are identified by all provided keys
// If you provide, for example, three keys to this function, all items returned can be identified
// by all those keys. If no results are returned, it means that no item stored in the storage can be
// identified by all keys at the same time.
func (memoryStorage *MemoryStorage) Search(keys []string) []Item {
	var mutex sync.Mutex
	keys = getLowercaseKeys(keys)
	itemOccurrencesCounter := make(map[*Item]int, 0)
	var wg sync.WaitGroup
	for _, key := range keys {
		wg.Add(1)
		go func(key string) {
			items := memoryStorage.Get(key)

			for _, item := range items {
				mutex.Lock()
				if numberOccurrences, itemFound := itemOccurrencesCounter[item]; itemFound {
					itemOccurrencesCounter[item] = numberOccurrences + 1
				} else {
					itemOccurrencesCounter[item] = 1
				}
				mutex.Unlock()
			}

			wg.Done()
		}(key)
	}

	wg.Wait()

	searchResults := make([]Item, 0)

	for item, numberOccurrences := range itemOccurrencesCounter {
		if numberOccurrences == len(keys) {
			searchResults = append(searchResults, *item)
		}
	}

	return searchResults
}

func getLowercaseKeys(keys []string) []string {
	lowercaseKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		lowercaseKeys = append(lowercaseKeys, strings.ToLower(key))
	}

	return lowercaseKeys
}
