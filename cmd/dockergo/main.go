package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/minhajuddinkhan/dockergo"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(dockergo.SayHello()))
	})
	fmt.Println("===============server starting==============!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
