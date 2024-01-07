package application

import (
	"fmt"
	"net/http"
)

func MakeShort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("make short called")
}

func MakeLong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("make long called")
}
