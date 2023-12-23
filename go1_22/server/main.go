package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /test/{x}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
		x := r.PathValue("x")
		fmt.Println(x)
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port 8080")
}
