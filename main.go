package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res := &response{Message: "42 - The Answer to the Ultimate Question of Life, The Universe, and Everything."}

		count := time.Now().Second()
		gen := random(res)

		for i := 0; i < count; i++ {
			gen()
		}

		out, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(out))

	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string `json:"message"`
	Random  []int  `json:"random"`
}

func random(res *response) func() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() {
		res.Random = append(res.Random, r.Intn(42))
	}
}
