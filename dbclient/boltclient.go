package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/callistaenterprise/goblog/accountservice/model"
	"log"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open( "accounts.db", 0600, nil )

	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) QueryAccount( accountId string ) (model.Account, error) {

	account := model.Account{}

	e := bc.boltDB.View( func(tx *bolt.Tx) error {

		b := tx.Bucket( []byte("AccountBucket"))

		accountBytes := b.Get([]byte(accountId))

		if accountBytes == nil {
			return fmt.Errorf( "No account found for " + accountId )
		}

		err := json.Unmarshal( accountBytes, &account )

		if err != nil {
			return fmt.Errorf( "QueryAccount() error: " + err.Error() )
		}

		return nil
	})

	if e != nil {
		return model.Account{}, e
	}

	return account, nil
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initializeBucket() {

	e := bc.boltDB.Update( func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket( []byte("AccountBucket") )

		if err != nil {
			return fmt.Errorf( "Create bucket failed: %s", err )
		}

		return nil
	})

	if e != nil {
		fmt.Println( "initializeBucket Error: " + e.Error() )
	}
}

func (bc *BoltClient) seedAccounts() {

	total := 100

	for i := 0 ; i < total ; i++ {

		key := strconv.Itoa( 10000 + i )

		acc := model.Account {
			Id: key,
			Name: "Person_" + strconv.Itoa(i),
		}

		jsonBytes, _ := json.Marshal(acc)

		e := bc.boltDB.Update( func( tx *bolt.Tx) error {
			b := tx.Bucket( []byte("AccountBucket"))
			err := b.Put( []byte(key), jsonBytes )
			return err
		})

		if e != nil {
			fmt.Printf( "seedAccounts() Error: %s", e.Error() )
		}
	}

	fmt.Printf( "Seeded %v fake accounts...\n", total )
}
