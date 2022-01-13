package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getCoinsName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func main() {

	resp, err := http.Get("https://api.coinranking.com/v2/coins")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/coinsName", getCoinsName)
}
