package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/godoc-os":   "https://golang.org/pkg/os/ ",
		"/godoc-http": "https://golang.org/pkg/net/http",
	}

	// json :=
	// [
	// 	{"path" : "/urlshort" , "url" : " "},
	// 	{"path" : "/urlshort" , "url" : " "},
	// 	{"path" : "/urlshort" , "url" : " "},
	// ]

	yaml := `
 - path: /urlshort
   url: https://github.com/gophercises/urlshort
 - path: /urlshort-final
   url: https://github.com/gophercises/urlshort/tree/solution
 `

	mHandler := MapHandler(pathsToUrls, mux)
	yHandler, err := YAMLHandler([]byte(yaml), mHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on: 8080")
	http.ListenAndServe(":8080", yHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fallback)
	return mux
}

func fallback(wri http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wri, "Link not found !!!")
}
