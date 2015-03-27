package pezsearch

import "strconv"

type DB interface {
	find(params *queryParams) (count int, results []*Record)
}

// Record is a DB record
type Record struct {
	ID      string      `json:"_id"`
	Type    string      `json:"_type"`
	URI     string      `json:"_uri"`
	Content interface{} `json:"_content"`
}

// Item - data for the item resource
type Item struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeID      string `json:"type-id"`
}

// Type - data for the type resource
type Type struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"status"`
}

// fake db
type fakeDB struct {
	Data []*Record
}

// fake find logic: there is no full text search implemented
// scope:item yields 10 records
// scope:type yields 2 records
// no scope returns all 12 records
// will honor limit and offset
func (db *fakeDB) find(params *queryParams) (count int, results []*Record) {

	if params.scope == TypeResource {
		results = db.Data[:2]
	} else if params.scope == ItemResource {
		results = db.Data[2:]
	} else {
		results = db.Data
	}
	count = len(results)

	//TODO(dnem): check for limit and offset and return accordingly
	return
}

func initDB() DB {
	nbr := 10
	fdb := &fakeDB{
		Data: make([]*Record, 0, nbr+2),
	}

	//Fill it up
	//Types
	fdb.Data = append(fdb.Data, &Record{
		ID:   "100",
		Type: TypeResource,
		URI:  "http://inventory.pez.io/v1/types/100",
		Content: Type{
			ID:          "type-id-100",
			Name:        "Type 100",
			Description: "Description for Type 100",
			Available:   true,
		},
	})

	fdb.Data = append(fdb.Data, &Record{
		ID:   "200",
		Type: TypeResource,
		URI:  "http://inventory.pez.io/v1/types/200",
		Content: Type{
			ID:          "type-id-200",
			Name:        "Type 200",
			Description: "Description for Type 200",
			Available:   true,
		},
	})

	for i := 1; i <= nbr; i++ {
		a := strconv.Itoa(i)
		x := 1
		if i%2 == 0 {
			x = 2
		}
		itm := Item{
			ID:          "item-id-" + a,
			Name:        "Item " + a,
			Description: "Description for Item " + a,
			TypeID:      "type-id-" + strconv.Itoa(x) + "00",
		}
		fdb.Data = append(fdb.Data, &Record{
			ID:      a,
			Type:    ItemResource,
			URI:     "http://inventory.pez.io/v1/items/" + a,
			Content: itm,
		})
	}
	return DB(fdb)
}
