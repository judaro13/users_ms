package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Name        string `json:"Name,omitempty"`
	Email       string `json:"Email,omitempty"`
	Password    string `json:"Password,omitempty"`
	Validated   bool   `json:"Validated,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Country     string `json:"Country,omitempty"`
	City        string `json:"City,omitempty"`
	Address     string `json:"Address,omitempty"`
}

var (
	ErrNameRequired     = errors.New("no key 'Name' gived")
	ErrEmailRequired    = errors.New("no key 'Email' gived")
	ErrPasswordRequired = errors.New("no key 'password' gived")
	ErrPhoneRequired    = errors.New("no key 'phoneNumber' gived")

	DBHost     = os.Getenv("DB_HOST")
	DBPort     = 5432
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
)

func ValidateUser(user *User) error {
	if len(user.Name) == 0 {
		return ErrNameRequired
	}
	if len(user.Email) == 0 {
		return ErrEmailRequired
	}
	if len(user.Password) == 0 {
		return ErrPasswordRequired
	}
	if len(user.PhoneNumber) == 0 {
		return ErrPhoneRequired
	}
	return nil
}

func Save(user *User) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// I didnt know how to do this pretty :(
	// GRANT ALL PRIVILEGES ON TABLE users TO Dbuser;
	// GRANT ALL PRIVILEGES ON TABLE users_uid_seq TO Dbuser;
	var lastInsertId int
	err = db.QueryRow("INSERT INTO users(name,email,password,validated,phone_number,country,city,address) VALUES($1,$2,$3,$4,$5,$6,$7,$8) returning uid;", user.Name, user.Email, user.Password, user.Validated, user.PhoneNumber, user.Country, user.City, user.Address).Scan(&lastInsertId)

	if err != nil {
		return err
	}

	fmt.Println("last inserted id =", lastInsertId)
	return nil
}

func Store(userJson []byte) {

	user := new(User)
	err := json.Unmarshal(userJson, user)

	if err != nil {
		fmt.Printf("%s: %s", err.Error(), userJson)
		return
	}

	if err := ValidateUser(user); err != nil {
		fmt.Printf("%s: %s", err.Error(), userJson)
		return
	}

	if err := Save(user); err != nil {
		fmt.Printf("%s: %s", err.Error(), userJson)
		return
	}
}
