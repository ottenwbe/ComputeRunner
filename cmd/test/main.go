package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

func main() {
	// Create a sample struct
	type Account struct {
		Name string
		ID   uuid.UUID
	}

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"account": &memdb.TableSchema{
				Name: "account",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert some people
	people := []*Account{
		&Account{"joe@aol.com", uuid.New()},
		&Account{"lucy@aol.com", uuid.New()},
		&Account{"tariq@aol.com", uuid.New()},
		&Account{"dorothy@aol.com", uuid.New()},
	}
	for _, p := range people {
		if err := txn.Insert("account", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// Lookup by email
	raw, err := txn.First("account", "id", "joe@aol.com")
	if err != nil {
		panic(err)
	}

	// Say hi!
	fmt.Printf("Hello %s .. %s!\n", raw.(*Account).Name, raw.(*Account).ID)

	// List all the people
	it, err := txn.Get("account", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the people:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Account)
		fmt.Printf("  %s\n", p.Name)
	}

	//// Range scan over people with ages between 25 and 35 inclusive
	//it, err = txn.LowerBound("account", "age", 25)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("People aged 25 - 35:")
	//for obj := it.Next(); obj != nil; obj = it.Next() {
	//	p := obj.(*Account)
	//	if p.Age > 35 {
	//		break
	//	}
	//	fmt.Printf("  %s is aged %d\n", p.Name, p.Age)
	//}
	// Output:
	// Hello Joe!
	// All the people:
	//   Dorothy
	//   Joe
	//   Lucy
	//   Tariq
	// People aged 25 - 35:
	//   Joe is aged 30
	//   Lucy is aged 35
}
