package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	mrand "math/rand"
	"net/http"
	"os"
	"time"

	"syreclabs.com/go/faker"
)

type User struct {
	Name        string
	Email       string
	PhoneNumber string
	IP          string
	CardNumber  string
	SSN         string
	AccessToken string
}

var ssnList []string

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		var users []User

		for i := 0; i < 10; i++ {
			randAccKey := make([]byte, 10)
			_, err := rand.Read(randAccKey)
			if err != nil {
				panic(err)
			}

			user := User{
				Name:        faker.Name().Name(),
				Email:       faker.Internet().Email(),
				PhoneNumber: faker.PhoneNumber().CellPhone(),
				IP:          faker.Internet().IpV4Address(),
				CardNumber:  faker.Finance().CreditCard(faker.CC_VISA),
				SSN:         pickSSN(),
				AccessToken: fmt.Sprintf("%x", randAccKey),
			}

			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		err := encoder.Encode(users)
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func pickSSN() string {
	if ssnList == nil {
		file, err := os.Open("./ssn.txt")
		if err != nil {
			panic("Can't load file with ssn numbers")
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ssnList = append(ssnList, string(line))
		}
	}

	mrand.Seed(time.Now().UnixNano())
	return ssnList[mrand.Intn(len(ssnList))]
}
