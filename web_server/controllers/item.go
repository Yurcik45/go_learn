package controllers

import (
  "fmt"
  "net/http"
)

type UserData struct {
	age int16
	id int64
	name string
}

func get_item(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func post_item(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}

func patch_item(w http.ResponseWriter, r *http.Request) {
	fmt.Println("patch")
}

func delete_item(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
}

func Item_handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    case "GET": 	 get_item(w, r)
    case "POST": 	 post_item(w, r)
		case "PATCH":  patch_item(w, r)
		case "DELETE": delete_item(w, r)
    default: fmt.Println("method ", r.Method, " doesn't support")
  }
}
