package main

import (
	"fmt"
	"net/http"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
		hi
		</head>
		<body>
			<h1>
			123
			</h1>
		</body>
	</html>
`)
}

func main() {
	http.ListenAndServe("localhost:8080", http.HandleFunc(myfunc))
}
