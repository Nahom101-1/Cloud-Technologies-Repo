package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"
	response, err := http.Get(url)
	if err != nil {
		//error and then OS.exit()
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}
		decoder := json.NewDecoder(response.Body)
		/*decoder.DisallowUnknownFields()*/ //prevents unknown fields in JSON
		err := decoder.Decode(&todoItem)    // decode the JSON into the struct
		if err != nil {
			log.Fatal(err)
		}
		/*fmt.Printf("Data from API: %+v\n", todoItem)*/
		//COvert back to JSON
		todo, _ := json.MarshalIndent(todoItem, "", "\t")

		fmt.Println(string((todo)))

	}
}
