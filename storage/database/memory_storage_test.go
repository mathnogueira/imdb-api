package database_test

import (
	"strings"

	"github.com/mathnogueira/imdb-api/storage/database"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("MemoryStorage", func() {

	var storage database.Storage
	var users []userItem
	logger := zap.NewNop()

	BeforeEach(func() {
		storage = database.NewMemoryStorage(logger)
		users = []userItem{
			{
				user{Name: "John Doe"},
			},
			{
				user{Name: "Jane Doe"},
			},
			{
				user{Name: "John Cena"},
			},
		}
	})

	It("Should store keywords and be able to return them in a search", func() {
		for _, user := range users {
			storage.Save(user)
		}

		listOfUsersWithSurnameDoe := storage.Get("Doe")

		Expect(listOfUsersWithSurnameDoe).To(HaveLen(2))
	})

	It("Should not duplicate items in the same bucket in case of duplicated keys", func() {
		users = append(users, userItem{
			user{Name: "Obu Obu Obu"}, // Real name, btw: https://www.youtube.com/watch?v=__bAh8nV9MI
		})
		for _, user := range users {
			storage.Save(user)
		}

		listOfUsersNamedObu := storage.Get("Obu")

		Expect(listOfUsersNamedObu).To(HaveLen(1))
	})

	It("Should return items that have all provided keys", func() {
		for _, user := range users {
			storage.Save(user)
		}

		items := storage.Search([]string{"John", "Cena"})

		Expect(items).To(HaveLen(1))

		johnCenaUser := items[0].GetContent().(user)

		Expect(johnCenaUser.Name).To(Equal("John Cena"))
	})

	It("lowercase and uppercase keys should return the same results", func() {
		for _, user := range users {
			storage.Save(user)
		}

		itemsFromLowercaseKeys := storage.Search([]string{"john", "cena"})
		itemsFromUppercaseKeys := storage.Search([]string{"JOHN", "CENA"})

		Expect(itemsFromLowercaseKeys).To(HaveLen(1))
		Expect(itemsFromUppercaseKeys).To(HaveLen(1))
	})
})

type user struct {
	Name string
}

type userItem struct {
	Content user
}

func (userItem userItem) GetContent() interface{} {
	return userItem.Content
}

func (userItem userItem) GetKeys() []string {
	return strings.Split(userItem.Content.Name, " ")
}
