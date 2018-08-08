package main

import (
	"log"
	"net/http"

	"github.com/elojah/crowl/pkg/twitter"
)

func main() {
	s := twitter.NewService()
	_ = s
	log.Fatal(http.ListenAndServe(":8080", nil))
}
