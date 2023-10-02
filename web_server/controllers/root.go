package controllers

import (
	"fmt"
	"net/http"
)

func Root_handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root handle")
  accepted := r.Method == "GET"   ||
              r.Method == "POST"  ||
		          r.Method == "PATCH" ||
              r.Method == "DELETE"

	if !accepted {
		fmt.Println("method ", r.Method, " doesn't support")
		http.Error(w, "Method is not supported.", http.StatusNotFound)
    return
	}

	fmt.Println("root direction")
}
