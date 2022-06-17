package account

import (
	"errors"

	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

// Ledger of all accounts. Facade to any db that stores accounts.
type Ledger struct {
	db *memdb.MemDB
}

// Add an account to the ledger
func (l *Ledger) Add(account *Account) error {
	if account == nil {
		return NILERROR
	}

	if l.doesExists(account.Name) {
		logrus.Infof("Doublicate : %v", account.Name)
		return ALREADYEXISTSERROR
	}

	// Create a write transaction
	txn := l.db.Txn(true)
	
	err := txn.Insert(accountTable, account)
	if err != nil {
		logrus.WithField("id", account.ID).Error("Could not write an account", err)
		txn.Abort()
		return err //Could also be a "wrapped" error dedicated for our data type and not depending on other packages
	}

	txn.Commit()
	return nil
}

func (l *Ledger) doesExists(name string) bool {
	logrus.Infof("Test if it exists: %v", name)
	return l.RetrieveByName(name) != nil
}

// RetrieveByName an account from the ledger
func (l *Ledger) RetrieveByName(name string) (account *Account) {
	// Create read-only transaction
	txn := l.db.Txn(false)
	defer txn.Abort()

	account = nil
	raw, err := txn.First(accountTable, nameIndex, name)
	if err != nil {
		logrus.WithField(nameIndex, name).Error("Error retrieving account by name", err)
	} else if raw == nil {
		logrus.WithField(nameIndex, name).Infof("Could not find account by name %v", name)
	} else {
		account = raw.(*Account)
	}

	return account
}

// RetrieveByID an account from the ledger
func (l *Ledger) RetrieveByID(id string) (account *Account) {
	// Create read-only transaction
	txn := l.db.Txn(false)
	defer txn.Abort()

	account = nil
	raw, err := txn.First(accountTable, idIndex, id)
	if err != nil {
		logrus.WithField(idIndex, id).Error("Error retrieving account by id", err)
	} else if raw == nil {
		logrus.WithField(idIndex, id).Infof("Could not find account by id %v", id)
	} else {
		account = raw.(*Account)
	}

	return account
}

// Accounts are stored here locally
var Accounts *Ledger

// NewAccountLedger inits the accounts db
func NewAccountLedger() *Ledger {

	// Create the Accounts schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			accountTable: &memdb.TableSchema{
				Name: accountTable,
				Indexes: map[string]*memdb.IndexSchema{
					idIndex: &memdb.IndexSchema{
						Name:    idIndex,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					nameIndex: &memdb.IndexSchema{
						Name:    nameIndex,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	// Create a new accounts data base
	accounts, err := memdb.NewMemDB(schema)
	if err != nil {
		logrus.Error("Could not create new MemDB", err)
		panic(err)
	}

	return &Ledger{
		db: accounts,
	}
}

func init() {
	Accounts = NewAccountLedger()
	logrus.Info("Accounts DB created")
}

var (
	NILERROR           = errors.New("nil account")
	ALREADYEXISTSERROR = errors.New("account already exists")
)

const (
	accountTable = "account"
	idIndex      = "id"
	nameIndex    = "name"
)
