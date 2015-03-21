package pezsearch

type item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeID      string `json:"type-id"`
}

func listItems() (i []item) {
	i1 := &item{
		ID:          "1",
		Name:        "Item 1",
		Description: "Item 1 Description",
		TypeID:      "abc",
	}
	i2 := &item{
		ID:          "2",
		Name:        "Item 2",
		Description: "Item 2 Description",
		TypeID:      "def",
	}
	i = []item{*i1, *i2}
	return
}

func getItem(id string) (i *item) {
	i = &item{
		ID:          id,
		Name:        "Item " + id,
		Description: "Item " + id + " Description",
		TypeID:      "ghi",
	}
	return
}

func listItemsByType(id string) (i []item) {
	i1 := &item{
		ID:          "4",
		Name:        "Item 1",
		Description: "Item 1 Description",
		TypeID:      id,
	}
	i2 := &item{
		ID:          "5",
		Name:        "Item 2",
		Description: "Item 2 Description",
		TypeID:      id,
	}
	i = []item{*i1, *i2}
	return
}
