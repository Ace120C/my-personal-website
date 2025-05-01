package main

import (
	"net/http"
	"fmt"
)


func main()  {
	fs := http.FileServer(http.Dir("../my-personal-website/"))
	http.Handle("/", fs)

	fmt.Println("Running Server at Port :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
