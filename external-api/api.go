package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
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
}

var ssnList []string

func main() {
	var users []User

	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		for i := 0; i < 10; i++ {
			user := User{
				Name:        faker.Name().Name(),
				Email:       faker.Internet().Email(),
				PhoneNumber: faker.PhoneNumber().CellPhone(),
				IP:          faker.Internet().IpV4Address(),
				CardNumber:  faker.Finance().CreditCard(faker.CC_VISA),
				SSN:         pickSSN(),
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

	rand.Seed(time.Now().Unix())
	return ssnList[rand.Intn(len(ssnList))]
}
