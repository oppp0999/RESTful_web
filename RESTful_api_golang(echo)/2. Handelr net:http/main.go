package main

import (
	"fmt"
	"net/http"
)

/*type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}*/

type MyWebserverType bool

func (m MyWebserverType) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "well, hello")
	fmt.Fprintf(w, "requst is : ", r) //응답 작성기
}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
