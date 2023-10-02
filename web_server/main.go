package main

import (
	"fmt"
	"log"
	"net/http"

	"my_module/controllers"
)

func my_server() {
	http.HandleFunc("/", controllers.Root_handle)
	http.HandleFunc("/item", controllers.Item_handle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} else {
	  fmt.Println("started on port 8080")
  }
}

func main() {
	my_server()
}
