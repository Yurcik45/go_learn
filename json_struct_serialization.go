package main

import (
	"fmt"
	"encoding/json"
)

type SerializablePerson struct {
	Id int `json:"person_id"`
	Name string `json:"person_name"`
}

func main() {
	 person := SerializablePerson {
		Id: 119,
		Name: "Gogi",
	 }

	data, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	fmt.Println(string(data))

	var decodedPerson SerializablePerson
    err = json.Unmarshal(data, &decodedPerson)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("%+v\n", decodedPerson)
}
