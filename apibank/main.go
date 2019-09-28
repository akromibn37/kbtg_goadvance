package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type mystruct struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {
	answer := mystruct{}
	getJson("https://jsonplaceholder.typicode.com/users/1")
	getJsonx("https://jsonplaceholder.typicode.com/users/1", &answer)
	fmt.Println("Answer :", answer)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// func getJson(url string, target interface{}) {
func getJson(url string) {
	r, err := myClient.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var data mystruct
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	// os.Exit(0)

	// return json.NewDecoder(r.Body).Decode(target)
}

func getJsonx(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
