package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jamisonwilliams99/encrypted-password-safe/caesar"
)

var passwordsBucket = []byte("passwords")

var db *bolt.DB

// need to modify the operation of the db in order for the UsedFor field to work
type Password struct {
	Id      int    `json:"id"`
	Value   string `json:"value"`
	UsedFor string `json:"usefor"`
}

// encodes the password into JSON format to be written to the database
func encodePassword(id int, encryptedPassword string, usedFor string) ([]byte, error) {
	pw := Password{
		Id:      id,
		Value:   encryptedPassword,
		UsedFor: usedFor,
	}
	jsonPw, err := json.Marshal(pw) // encode password struct in json so that it can be stored in a single key-value pair in the database
	if err != nil {
		return nil, err
	}
	return jsonPw, nil
}

// decodes a JSON byte slice into a Password object (returns reference to object)
func decodePassword(jsonPw []byte) (*Password, error) {
	var pw Password
	err := json.Unmarshal(jsonPw, &pw)
	if err != nil {
		return nil, err
	}
	return &pw, nil
}

// converts interger to 8 byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// converts byte slice to integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(passwordsBucket)
		return err
	})
}

func CreatePassword(password string, key string, usedFor string) (string, error) {
	var retId int
	encryptedPassword := caesar.Encrypt(password, key)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		id64, _ := b.NextSequence()
		retId = int(id64)
		id := itob(retId)

		jsonPw, err := encodePassword(int(id64), encryptedPassword, usedFor)
		if err != nil {
			return err
		}

		return b.Put(id, jsonPw)
	})
	if err != nil {
		return "", err
	}
	return encryptedPassword, nil
}

// will decrypt the password specified by the id with the provided key, regardless of if the key is correct
func RetrievePassword(id int, key string) (string, string, error) {
	var jsonPw []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		jsonPw = b.Get(itob(id))
		return nil
	})
	if err != nil {
		return "", "", err
	}
	pw, err := decodePassword(jsonPw)
	if err != nil {
		return "", "", err
	}
	encryptedPassword := pw.Value
	return caesar.Decrypt(encryptedPassword, key), encryptedPassword, nil
}

// TODO - make sure valid id in caller
func UpdatePassword(id int, newPassword string, key string) error {
	var jsonPw []byte
	encryptedPassword := caesar.Encrypt(newPassword, key)

	// retrieve current password from the db to extract the UsedFor attribute
	// - should possibly make the retrieving the JSON password routine a function
	//   since it is used here and in the UpdatePassword function
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		jsonPw = b.Get(itob(id))
		return nil
	})
	if err != nil {
		return err
	}
	pw, err := decodePassword(jsonPw)
	usedFor := pw.UsedFor

	// write new password to database
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		jsonPw, err := encodePassword(id, encryptedPassword, usedFor)
		if err != nil {
			return err
		}
		return b.Put(itob(id), jsonPw)
	})
}

// TODO - reorder ids after deletion
func DeletePassword(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		return b.Delete(itob(id))
	})
}

// returns a list of every password id and what the password is used for
func AllPasswords() ([]Password, error) {
	var passwords []Password
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordsBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			pw, _ := decodePassword(v) // might want to handle this error later
			passwords = append(passwords, *pw)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return passwords, nil
}
