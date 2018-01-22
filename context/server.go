package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", LazyServer)
	http.ListenAndServe(":1111", nil)
}

func LazyServer(w http.ResponseWriter, r *http.Request) {
	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "Go! slow %v\n", headOrTails)
		fmt.Printf("Go! slow %v\n", headOrTails)
		return
	}

	fmt.Fprintf(w, "Go! quick %v\n", headOrTails)
	fmt.Printf("Go! quick %v\n", headOrTails)
	return
}
